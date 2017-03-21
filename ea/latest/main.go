package main

import (
	"github.com/rainchasers/report"
	"os"
	"strconv"
	"time"
)

const (
	maxPublishPerSecond  = 20
	httpTimeoutInSeconds = 30
	httpUserAgent        = "Rainchaser Bot <hello@rainchasers.com>"
)

// Responds to environment variables:
//   UPDATE_EVERY_X_SECONDS (default 15*60)
//   SHUTDOWN_AFTER_X_SECONDS (default 24*60*60)
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
	// setup telemetry and golang defaults
	report.StdOut()
	report.Global(report.Data{"service": "ea.latest", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	report.RuntimeStatsEvery(30 * time.Second)
	defer report.Drain()

	// parse env vars
	updatePeriodSeconds, err := strconv.Atoi(os.Getenv("UPDATE_EVERY_X_SECONDS"))
	if err != nil {
		updatePeriodSeconds = 15 * 60
	}
	shutdownDeadline, err := strconv.Atoi(os.Getenv("SHUTDOWN_AFTER_X_SECONDS"))
	if err != nil {
		shutdownDeadline = 7 * 24 * 60 * 60
	}
	shutdownC := time.NewTimer(time.Second * time.Duration(shutdownDeadline)).C
	projectId := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")

	// decision on whether validating
	isValidating := projectId == ""
	var logs *LogBuffer
	if isValidating {
		logs = trackLogs()
	}
	report.Info("daemon.start", report.Data{
		"update_period":     updatePeriodSeconds,
		"shutdown_deadline": shutdownDeadline,
		"project_id":        projectId,
		"pubsub_topic":      topicName,
	})

	// discover EA gauging stations
	refSnapshots, err := discover()
	if err != nil {
		report.Action("discovered.fail", report.Data{"error": err.Error()})
		return err
	}
	report.Info("discovered.ok", report.Data{"count": len(refSnapshots)})

	// periodically get latest updates
	ticker := time.NewTicker(time.Second * time.Duration(updatePeriodSeconds))

updateLoop:
	for {
		tick := report.Tick()
		updates, err := update()
		if err != nil {
			report.Action("updated.fail", report.Data{"error": err.Error()})
			return err
		}
		report.Tock(tick, "updated.ok", report.Data{"count": len(updates)})

		if !isValidating {
			tick = report.Tick()
			err = publish(projectId, topicName, updates, refSnapshots)
			if err != nil {
				report.Action("published.fail", report.Data{"error": err.Error()})
				return err
			}
			report.Tock(tick, "published.ok", report.Data{"count": len(updates)})
		}

		select {
		case <-ticker.C:
		case <-shutdownC:
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
