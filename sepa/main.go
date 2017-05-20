package main

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/rainchasers/com.rainchasers.gauge/queue"
	"github.com/rainchasers/report"
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
	isValidating := projectID == ""
	topicName := os.Getenv("PUBSUB_TOPIC")

	// setup telemetry and logging
	log := report.New(report.Data{"service": "sepa", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	log.RuntimeStatEvery("runtime", 5*time.Minute)
	defer log.Stop()
	log.Info("daemon.start", report.Data{
		"update_period": updatePeriodSeconds,
		"timeout":       timeout,
		"project_id":    projectID,
		"pubsub_topic":  topicName,
	})

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
			log.Info("daemon.interrupt", report.Data{})
			shutdown()
		case <-ctx.Done():
		}
	}()

	// discover SEPA gauging stations
	stations, err := discover()
	if err != nil {
		<-log.Action("discovered.failed", report.Data{"error": err.Error()})
		return err
	}
	if isValidating {
		stations = stations[0:5]
	}
	log.Info("discovered.ok", report.Data{"count": len(stations)})

	// calculate tick rate and spawn individual gauge download CSVs
	tickerMs := updatePeriodSeconds * 1000 / len(stations)
	minTickerMs := 1000 / maxDownloadPerSecond
	if tickerMs < minTickerMs {
		tickerMs = minTickerMs
	}
	n := 0
	ticker := time.NewTicker(time.Millisecond * time.Duration(tickerMs))
	defer ticker.Stop()

	// open connection to pubsub
	topic, err := queue.New(ctx, projectID, topicName)
	if err != nil {
		return err
	}
	defer topic.Stop()
	nPubErr := 0

updateLoop:
	for {
		i := n % len(stations)

		tick := log.Tick()
		readings, err := getReadings(stations[i].DataURL)
		if err != nil {
			log.Tock(tick, "updated.fail", report.Data{
				"url":   stations[i].DataURL,
				"error": err.Error(),
			})
		} else {
			log.Tock(tick, "updated.ok", report.Data{
				"url":   stations[i].DataURL,
				"count": len(readings),
			})
		}

		err = topic.Publish(context.Background(), &gauge.Snapshot{
			Station:  stations[i],
			Readings: readings,
		})
		if err != nil {
			<-log.Action("publish.fail", report.Data{
				"url":   stations[i].DataURL,
				"error": err.Error(),
			})
			nPubErr++
		}
		if nPubErr > 100 {
			shutdown()
		}

		n++
		select {
		case <-ticker.C:
		case <-ctx.Done():
			break updateLoop
		}
	}

	// validate log stream on shutdown if required
	if isValidating {
		if log.Count("discovered.ok") != 1 {
			return errors.New("discovered.ok event expected but not present")
		}
		if log.Count("updated.ok") < 1 {
			return errors.New("updated.ok event expected but not present")
		}
	}

	return nil
}
