package main

import (
	"context"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/queue"
	"github.com/rainchasers/report"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	maxDownloadPerSecond = 1
)

// Responds to environment variables:
//   UPDATE_EVERY_X_SECONDS (default 15*60)
//   SHUTDOWN_AFTER_X_SECONDS (default 7*24*60*60)
//   PROJECT_ID (no default, blank for validation mode)
//   PUBSUB_TOPIC (no default, blank for validation mode)
//
func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	// parse env vars
	updatePeriodSeconds, err := strconv.Atoi(os.Getenv("UPDATE_EVERY_X_SECONDS"))
	if err != nil {
		updatePeriodSeconds = 15 * 60
	}
	timeout, err := strconv.Atoi(os.Getenv("SHUTDOWN_AFTER_X_SECONDS"))
	if err != nil {
		timeout = 7 * 24 * 60 * 60
	}
	projectID := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")

	// setup telemetry and logging
	report.StdOut()
	report.Global(report.Data{"service": "sepa", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	report.RuntimeStatsEvery(30 * time.Second)
	defer report.Drain()
	report.Info("daemon.start", report.Data{
		"update_period": updatePeriodSeconds,
		"timeout":       timeout,
		"project_id":    projectID,
		"pubsub_topic":  topicName,
	})

	// decision on whether validating logs
	isValidating := projectID == ""
	var logs *LogBuffer
	if isValidating {
		logs = trackLogs()
	}

	// create daemon context
	ctx, shutdown := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	defer shutdown()
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		select {
		case <-sigC:
			report.Info("daemon.interrupt", report.Data{})
			shutdown()
		case <-ctx.Done():
		}
	}()

	// discover SEPA gauging stations
	stations, err := discover()
	if err != nil {
		report.Action("discovered.failed", report.Data{"error": err.Error()})
		return err
	}
	if isValidating {
		stations = stations[0:5]
	}
	report.Info("discovered.ok", report.Data{"count": len(stations)})

	// calculate tick rate and spawn individual gauge download CSVs
	tickerMs := updatePeriodSeconds * 1000 / len(stations)
	minTickerMs := 1000 / maxDownloadPerSecond
	if tickerMs < minTickerMs {
		tickerMs = minTickerMs
	}
	n := 0
	ticker := time.NewTicker(time.Millisecond * time.Duration(tickerMs))
	defer ticker.Stop()

	// open connection to publish topic
	var topic *queue.Topic
	nPubErr := 0
	if !isValidating {
		topic, err = queue.New(ctx, projectID, topicName)
		if err != nil {
			return err
		} else {
			defer topic.Stop()
		}
	}

updateLoop:
	for {
		i := n % len(stations)

		tick := report.Tick()
		readings, err := getReadings(stations[i].DataURL)
		if err != nil {
			report.Tock(tick, "updated.fail", report.Data{
				"url":   stations[i].DataURL,
				"error": err.Error(),
			})
		} else {
			report.Tock(tick, "updated.ok", report.Data{
				"url":   stations[i].DataURL,
				"count": len(readings),
			})
		}

		if !isValidating {
			err = topic.Publish(context.Background(), &gauge.Snapshot{
				Station:  stations[i],
				Readings: readings,
			})
			if err != nil {
				report.Action("publish.fail", report.Data{
					"url":   stations[i].DataURL,
					"error": err.Error(),
				})
				nPubErr += 1
			}
		}
		if nPubErr > 100 {
			shutdown()
		}

		n += 1
		select {
		case <-ticker.C:
		case <-ctx.Done():
			break updateLoop
		}
	}

	// validate log stream on shutdown if required
	if isValidating {
		report.Drain()
		expect := map[string]int{
			"discovered.ok": VALIDATE_IS_PRESENT,
			"updated.ok":    VALIDATE_IS_PRESENT,
		}
		err := validateLogStream(logs, expect)
		if err != nil {
			return err
		}
	}

	return nil
}
