package main

import (
	"os"
	"strconv"
	"time"

	"github.com/robtuley/report"
)

// Responds to environment variables:
//   UPDATE_EVERY_X_SECONDS (default 15*60)
//   UPDATE_COUNT_BEFORE_SHUTDOWN (default 100)
//   PROJECT_ID (no default, blank for validation mode)
//   LATEST_PUBSUB_TOPIC (no default, blank for validation mode)
//   HISTORY_PUBSUB_TOPIC (no default, blank for validation mode)
//
func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	// setup telemetry and logging
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "sepa", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	report.RuntimeStatsEvery(30 * time.Second)

	// parse env vars
	updatePeriodSeconds, err := strconv.Atoi(os.Getenv("UPDATE_EVERY_X_SECONDS"))
	if err != nil {
		updatePeriodSeconds = 15 * 60
	}
	updateCountOnShutdown, err := strconv.Atoi(os.Getenv("UPDATE_COUNT_BEFORE_SHUTDOWN"))
	if err != nil {
		updateCountOnShutdown = 100
	}
	projectId := os.Getenv("PROJECT_ID")
	latestTopicName := os.Getenv("LATEST_PUBSUB_TOPIC")
	historyTopicName := os.Getenv("HISTORY_PUBSUB_TOPIC")

	// decision on whether validating logs
	isValidating := projectId == ""
	var validateC <-chan report.Data
	if isValidating {
		validateC = bufferLogStream(1000)
	}
	report.Info("daemon.start", report.Data{
		"update_period":            updatePeriodSeconds,
		"update_count_on_shutdown": updateCountOnShutdown,
		"project_id":               projectId,
		"latest_pubsub_topic":      latestTopicName,
		"history_pubsub_topic":     historyTopicName,
	})

	// discover SEPA gauging stations
	refSnapshots, err := discoverStations()
	if err != nil {
		report.Action("discovered.failed", report.Data{"error": err.Error()})
		return err
	}
	if isValidating {
		refSnapshots = refSnapshots[0:5]
	}
	report.Info("discovered.ok", report.Data{"count": len(refSnapshots)})

	// calculate tick rate and spawn individual gauge download CSVs
	tickerMillisecond := updatePeriodSeconds * 1000 / len(refSnapshots)
	if tickerMillisecond < 1000 {
		tickerMillisecond = 1000
	}
	nMax := updateCountOnShutdown * len(refSnapshots)
	n := 0
	ticker := time.NewTicker(time.Millisecond * time.Duration(tickerMillisecond))

updateTick:
	for range ticker.C {
		i := n % len(refSnapshots)

		tick := report.Tick()
		readings, err := getReadings(refSnapshots[i])
		if err != nil {
			report.Tock(tick, "updated.failed", report.Data{
				"url":   refSnapshots[i].DataURL,
				"error": err.Error(),
			})
		} else {
			report.Tock(tick, "updated.ok", report.Data{
				"url":   refSnapshots[i].DataURL,
				"count": len(readings),
			})
		}

		n = n + 1
		if n >= nMax {
			break updateTick
		}
	}
	ticker.Stop()

	// validate log stream on shutdown if required
	err = nil
	if isValidating {
		expect := map[string]int{
			"discovered.ok": VALIDATE_IS_PRESENT,
			"updated.ok":    VALIDATE_IS_PRESENT, // TODO: updateCountOnShutdown * len(refSnapshots),
		}
		time.Sleep(time.Second) // TODO: remove this! it allows log flush
		err = validateLogStream(validateC, expect)
	}
	return err
}
