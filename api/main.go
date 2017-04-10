package main

import (
	"context"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/queue"
	"github.com/rainchasers/report"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

// Responds to environment variables:
//   BOOTSTRAP_URL (no default)
//   SHUTDOWN_AFTER_X_SECONDS (default 7*24*60*60)
//   PROJECT_ID (no default)
//   PUBSUB_TOPIC (no default)
//   CONSUMER_NAME (no default)
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
	consumerName := os.Getenv("CONSUMER_NAME")
	timeout, err := strconv.Atoi(os.Getenv("SHUTDOWN_AFTER_X_SECONDS"))
	if err != nil {
		timeout = 7 * 24 * 60 * 60
	}

	// setup telemetry and logging
	report.StdOut()
	report.Global(report.Data{"service": "api", "daemon": consumerName + time.Now().Format("v2006-01-02-15-04-05")})
	report.RuntimeStatsEvery(30 * time.Second)
	defer report.Drain()
	report.Info("daemon.start", report.Data{
		"bootstrap_url": bootstrapURL,
		"project_id":    projectId,
		"pubsub_topic":  topicName,
		"consumer_name": consumerName,
		"timeout":       timeout,
	})

	// create daemon context
	ctx, shutdown := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer shutdown()

	// create gauge in-memory cache
	cache := gauge.NewCache(ctx, 7*24*time.Hour)

	// subscribe to gauge snapshot topic to populate gauge cache
	var counter uint64
	go func() {
		topic, err := queue.New(projectId, topicName)
		if err != nil {
			report.Action("topic.failed", report.Data{"error": err.Error()})
			shutdown()
			return
		}
		err = topic.Subscribe(ctx, consumerName, func(s *gauge.Snapshot, err error) {
			if err != nil {
				report.Action("msg.failed", report.Data{"error": err.Error()})
			} else {
				atomic.AddUint64(&counter, 1)
				cache.Add(s)
			}
		})
		if err != nil {
			report.Action("subscribe.failed", report.Data{"error": err.Error()})
			shutdown()
			return
		}
	}()

	// log cache status every 30s
	ticker := time.NewTicker(time.Second * 30)
logLoop:
	for {
		stat := cache.Stats()
		report.Info("cache.counts", report.Data{
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

	return nil
}
