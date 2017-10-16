package main

import (
	"context"
	"net/http"
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

// Environment Vars
// ================
//
// Internal API to bootstrap from:
//   COM_RAINCHASERS_API_INTERNAL_SERVICE_HOST (no default)
//   COM_RAINCHASERS_API_INTERNAL_SERVICE_PORT (no default)
//
// SSL cert configs:
//   SSL_KEY_FILE (no default)
//   SSL_CERT_FILE (no default)
//
// General config:
//   DAEMON_NAME (defaults to a timebased stamp to label the daemon)
//   SHUTDOWN_AFTER_X_SECONDS (default 7*24*60*60)
//
// Google PubSub subscription:
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
	var selfURL string
	host := os.Getenv("COM_RAINCHASERS_API_INTERNAL_SERVICE_HOST")
	port, err := strconv.Atoi(os.Getenv("COM_RAINCHASERS_API_INTERNAL_SERVICE_PORT"))
	if err == nil && len(host) > 0 {
		selfURL = "http://" + host + ":" + strconv.Itoa(port) + "/export"
	}
	daemonName := os.Getenv("DAEMON_NAME")
	if len(daemonName) == 0 {
		daemonName = time.Now().Format("v2006-01-02-15-04-05.9999")
	}
	sslKeyFilename := os.Getenv("SSL_KEY_FILE")
	sslCertFilename := os.Getenv("SSL_CERT_FILE")
	projectID := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")
	timeout, err := strconv.Atoi(os.Getenv("SHUTDOWN_AFTER_X_SECONDS"))
	if err != nil {
		timeout = 7 * 24 * 60 * 60
	}

	// telemetry and logging
	log := report.New(os.Stdout, report.Data{"service": "rc.api", "daemon": daemonName})
	log.RuntimeStatEvery("runtime", 5*time.Minute)
	defer log.Stop()
	log.Info("daemon.start", report.Data{
		"bootstrap_url": selfURL,
		"project_id":    projectID,
		"pubsub_topic":  topicName,
		"timeout":       timeout,
		"ssl_key_path":  sslKeyFilename,
		"ssl_cert_path": sslCertFilename,
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
	gaugeCache := gauge.NewCache(ctx, 36*time.Hour)

	// subscribe to gauge snapshot topic to populate gauge cache
	var counter uint64
	wg.Add(1)
	go func() {
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
				gaugeCache.Add(s)
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
			stat := gaugeCache.Stats()
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

	// setup request handler & perform startup actions
	h := &Handler{
		Log:           log,
		Gauge:         gaugeCache,
		IsReady:       false,
		ClientTimeout: 10 * time.Second,
	}
	go func() {
		c := bootstrapGaugeCache(selfURL, gaugeCache, log)
		<-c
		h.IsReady = true
	}()

	// HTTP server :8081 for k8s status checks + avro dumps
	mux8081 := http.NewServeMux()
	mux8081.HandleFunc("/k8s/", h.Readiness)
	mux8081.HandleFunc("/export", h.Export)
	server8081 := &http.Server{
		Addr:    ":8081",
		Handler: mux8081,
	}
	go func() {
		err := server8081.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			<-log.Action("http.error", report.Data{
				"port":  8081,
				"error": err.Error(),
			})
			shutdown()
			return
		}
	}()

	//  HTTP server :8080 for http public API
	//              :8443 for https public API
	mux := http.NewServeMux()
	mux.HandleFunc("/", h.API)
	server8080 := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	go func() {
		err := server8080.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			<-log.Action("http.error", report.Data{
				"port":  8080,
				"error": err.Error(),
			})
			shutdown()
			return
		}
	}()
	var server8443 *http.Server
	if sslKeyFilename != "" {
		server8443 = &http.Server{
			Addr:    ":8443",
			Handler: mux,
		}
		go func() {
			err := server8443.ListenAndServeTLS(sslCertFilename, sslKeyFilename)
			if err != nil && err != http.ErrServerClosed {
				<-log.Action("http.error", report.Data{
					"port":  8443,
					"error": err.Error(),
				})
				shutdown()
				return
			}
		}()
	}

	// On shutdown signal, wait for up to 20s for go-routines
	// and HTTP servers to close cleanly
	<-ctx.Done()
	h.IsReady = false
	tContext, tCancel := context.WithTimeout(context.Background(), 20*time.Second)
	wg.Add(1)
	go func() {
		if server8443 != nil {
			server8443.Shutdown(tContext)
		}
		server8080.Shutdown(tContext)
		server8081.Shutdown(tContext)
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
	case <-tContext.Done():
		<-log.Action("daemon.stopped.timeout", report.Data{})
	}
	tCancel()

	return nil
}
