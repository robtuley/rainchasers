package main

import (
	"context"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/queue"
	"github.com/rainchasers/report"
	"os"
	"strconv"
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
	// setup telemetry and logging
	report.StdOut()
	report.Global(report.Data{"service": "api", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	report.RuntimeStatsEvery(30 * time.Second)
	defer report.Drain()

	// parse env vars
	bootstrapURL := os.Getenv("BOOTSTRAP_URL")
	projectId := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")
	timeout, err := strconv.Atoi(os.Getenv("SHUTDOWN_AFTER_X_SECONDS"))
	if err != nil {
		timeout = 7 * 24 * 60 * 60
	}
	report.Info("daemon.start", report.Data{
		"bootstrap_url": bootstrapURL,
		"project_id":    projectId,
		"pubsub_topic":  topicName,
		"timeout":       timeout,
	})

	// create daemon context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer cancel()

	// create gauge in-memory cache
	cache := gauge.NewCache(ctx, 7*24*time.Hour)

	// subscribe to gauge snapshot topic to populate gauge cache
	topic, err := queue.New(projectId, topicName)
	if err != nil {
		report.Action("topic.failed", report.Data{"error": err.Error()})
		return err
	}
	err = topic.Subscribe("api", func(s *gauge.Snapshot, err error) {
		if err != nil {
			report.Action("msg.failed", report.Data{"error": err.Error()})
		} else {
			cache.Add(s)
		}
	})
	if err != nil {
		report.Action("subscribe.failed", report.Data{"error": err.Error()})
		return err
	}

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
		})

		select {
		case <-ticker.C:
		case <-ctx.Done():
			break logLoop
		}
	}
	ticker.Stop()

	return nil
}
