package main

import (
	"context"
	"fmt"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/queue"
	"github.com/rainchasers/report"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// Responds to environment variables:
//   BOOTSTRAP_URL (no default)
//   SHUTDOWN_AFTER_X_SECONDS (default 7*24*60*60)
//   PROJECT_ID (no default)
//   PUBSUB_TOPIC (no default)
//
func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	// parse env vars
	bootstrapURL := os.Getenv("BOOTSTRAP_URL")
	projectId := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")
	timeout, err := strconv.Atoi(os.Getenv("SHUTDOWN_AFTER_X_SECONDS"))
	if err != nil {
		timeout = 7 * 24 * 60 * 60
	}

	// setup telemetry and logging
	log := report.New(report.Data{"service": "api", "daemon": time.Now().Format("v2006-01-02-15-04-05.9999")})
	log.RuntimeStatEvery("runtime", 5*time.Minute)
	defer log.Stop()
	log.Info("daemon.start", report.Data{
		"bootstrap_url": bootstrapURL,
		"project_id":    projectId,
		"pubsub_topic":  topicName,
		"timeout":       timeout,
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
		select {
		case <-sigC:
			<-log.Info("daemon.interrupt", report.Data{})
			shutdown()
		case <-ctx.Done():
		}
		wg.Done()
	}()

	// create gauge in-memory cache
	cache := gauge.NewCache(ctx, 7*24*time.Hour)

	// subscribe to gauge snapshot topic to populate gauge cache
	var counter uint64
	wg.Add(1)
	go func() {
		topic, err := queue.New(ctx, projectId, topicName)
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
				cache.Add(s)
			}
		})
		topic.Stop()
		if err != nil {
			<-log.Action("subscribe.failed", report.Data{"error": err.Error()})
			shutdown()
			return
		}
		wg.Done()
	}()

	// log cache status every 30s
	wg.Add(1)
	go func() {
		ticker := time.NewTicker(time.Second * 30)
	logLoop:
		for {
			stat := cache.Stats()
			log.Info("cache.counts", report.Data{
				"station":         stat.StationCount,
				"all_reading":     stat.AllReadingCount,
				"max_reading":     stat.MaxReadingCount,
				"min_reading":     stat.MinReadingCount,
				"max_age_seconds": stat.OldestReading.Seconds(),
				"added":           atomic.LoadUint64(&counter),
			})
			atomic.StoreUint64(&counter, 0)

			select {
			case <-ticker.C:
			case <-ctx.Done():
				break logLoop
			}
		}
		ticker.Stop()
		wg.Done()
	}()

	// HTTP server :8081 for k8s status checks (readiness & liveness) and internal avro dumps
	started := time.Now()

	mux8081 := http.NewServeMux()
	mux8081.HandleFunc("/k8s/", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Now().Sub(started)
		if duration.Seconds() <= 10 {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		}
	})
	server8081 := &http.Server{
		Addr:    ":8081",
		Handler: mux8081,
	}
	go func() {
		if err := server8081.ListenAndServe(); err != nil {
			log.Action("http.error", report.Data{
				"port":  8081,
				"error": err.Error(),
			})
		}
	}()

	//  HTTP server :8080 for http public API
	//              :8443 for https public API

	// On shutdown signal, wait for up to 20s for go-routines & HTTP servers to close cleanly
	<-ctx.Done()
	terminationContext, _ := context.WithTimeout(context.Background(), 20*time.Second)
	wg.Add(1)
	go func() {
		server8081.Shutdown(terminationContext)
		wg.Done()
	}()
	c := make(chan bool)
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		<-log.Info("daemon.stopped", report.Data{})
	case <-terminationContext.Done():
		<-log.Action("daemon.stopped.timeout", report.Data{})
	}

	return nil
}
