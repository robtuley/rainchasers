package main

import (
	"context"
	"github.com/rainchasers/report"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

const (
	maxDownloadPerSecond = 1
	maxPublishPerSecond  = 20
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
	projectId := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")

	// setup telemetry and logging
	report.StdOut()
	report.Global(report.Data{"service": "sepa", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	report.RuntimeStatsEvery(30 * time.Second)
	defer report.Drain()
	report.Info("daemon.start", report.Data{
		"update_period": updatePeriodSeconds,
		"timeout":       timeout,
		"project_id":    projectId,
		"pubsub_topic":  topicName,
	})

	// decision on whether validating logs
	isValidating := projectId == ""
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

		n = n + 1
		select {
		case <-ticker.C:
		case <-ctx.Done():
			break updateLoop
		}
	}
	ticker.Stop()

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
