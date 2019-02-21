package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/queue"
	"github.com/rainchasers/report"
)

// Env:
//
//   DAEMON_NAME (defaults to a timebased stamp to label the daemon)
//   SHUTDOWN_AFTER_X_SECONDS (default 7*24*60*60)
//   PROJECT_ID (no default, blank indicates self-test mode)
//   PUBSUB_TOPIC (no default)
//   HONEYCOMB_API_KEY (no default, blank to skip honeycomb events)
//
func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	// parse env vars
	daemonName := os.Getenv("DAEMON_NAME")
	if len(daemonName) == 0 {
		daemonName = time.Now().Format("v2006-01-02-15-04-05.9999")
	}
	projectID := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")
	honeycombKey := os.Getenv("HONEYCOMB_API_KEY")
	timeout, err := strconv.Atoi(os.Getenv("SHUTDOWN_AFTER_X_SECONDS"))
	if err != nil {
		timeout = 7 * 24 * 60 * 60
	}
	if projectID == "" {
		timeout = 15
	}

	// telemetry and logging
	w := report.StdOutJSON()
	if len(honeycombKey) > 0 {
		w = w.And(report.Honeycomb(honeycombKey, "firestore"))
	}
	log := report.New(w, report.Data{"service": "firestore", "daemon": daemonName})
	log.RuntimeStatEvery("runtime", 5*time.Minute)
	defer log.Stop()
	log.Info("daemon.start", report.Data{
		"project_id":   projectID,
		"pubsub_topic": topicName,
		"timeout":      timeout,
	})

	// create daemon context
	var wg sync.WaitGroup
	ctx, shutdown := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer shutdown()
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		case <-sigC:
			<-log.Info("daemon.interrupt", report.Data{})
			shutdown()
		case <-ctx.Done():
		}
	}()

	// load river catalogue
	url := "https://app.rainchasers.com/catalogue.json"
	rivers, err := NewRiverCache(ctx, url)
	if err != nil {
		<-log.Action("daemon.stopped.rivercache", report.Data{
			"error": err.Error(),
		})
		shutdown()
		return err
	}
	log.Info("river.cache.created", report.Data{
		"count": rivers.Count(),
	})
	// TODO: update firebase rivers here

	// poll for river content updates
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(60 * time.Second)
	updateThenWait:
		for {
			select {
			case <-ticker.C:
			case <-ctx.Done():
				break updateThenWait
			}

			isChanged, err := rivers.Update()
			if err != nil {
				log.Action("river.cache.error", report.Data{"error": err.Error()})
				continue
			}
			if isChanged {
				log.Action("river.cache.changed", report.Data{
					"version": rivers.Version,
				})
				// TODO: update firebase here
			}
		}
		ticker.Stop()
	}()

	// subscribe to gauge snapshot topic to populate firebase
	var counter uint64
	wg.Add(1)
	go func() {
		defer wg.Done()

		if len(projectID) == 0 {
			// running in local test mode
			return
		}

		topic, err := queue.New(ctx, projectID, topicName)
		if err != nil {
			<-log.Action("topic.failed", report.Data{"error": err.Error()})
			shutdown()
			return
		}
		err = topic.Subscribe(ctx, "", func(s *gauge.Snapshot, err error) {
			if err != nil {
				log.Action("msg.failed", report.Data{"error": err.Error()})
			} else {
				atomic.AddUint64(&counter, 1)
				// do something with firebase here
			}
		})
		topic.Stop()
		if err != nil {
			<-log.Action("subscribe.failed", report.Data{"error": err.Error()})
			shutdown()
			return
		}
	}()

	// log updates received ok every 30s
	wg.Add(1)
	go func() {
		defer wg.Done()

		ticker := time.NewTicker(time.Second * 30)
	logLoop:
		for {
			select {
			case <-ticker.C:
			case <-ctx.Done():
				break logLoop
			}

			log.Info("received.ok", report.Data{
				"count": atomic.LoadUint64(&counter),
			})
			atomic.StoreUint64(&counter, 0)
		}
		ticker.Stop()
	}()

	// On shutdown signal, wait for up to 10s for go-routines to clean close
	<-ctx.Done()
	tContext, tCancel := context.WithTimeout(context.Background(), 10*time.Second)
	c := make(chan bool)
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		<-log.Info("daemon.stopped.ok", report.Data{})
	case <-tContext.Done():
		<-log.Action("daemon.stopped.timeout", report.Data{})
	}
	tCancel()

	// validate log stream on shutdown
	if log.Count("river.cache.created") != 1 {
		return errors.New("river.cache.created event expected but not present")
	}
	if err := log.LastError(); err != nil {
		return err
	}

	return nil
}
