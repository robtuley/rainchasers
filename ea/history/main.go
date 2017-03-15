package main

import (
	"github.com/robtuley/report"
	"os"
	"time"
)

const (
	maxPublishPerSecond  = 20
	httpTimeoutInSeconds = 300
	httpUserAgent        = "Rainchaser Bot <hello@rainchasers.com>"
)

// Responds to environment variables:
//   DATE (defaults to yesterday)
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
	var err error

	// setup telemetry and golang defaults
	report.StdOut()
	report.Global(report.Data{"service": "ea.history", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	report.RuntimeStatsEvery(1 * time.Second)

	// parse env vars
	requestedDay := os.Getenv("DATE")
	var day time.Time
	if len(requestedDay) == 0 {
		day = time.Now().AddDate(0, 0, -1)
	} else {
		day, err = time.Parse("2006-01-02", requestedDay)
		if err != nil {
			return err
		}
	}
	projectId := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")

	// decision on whether validating
	isValidating := projectId == ""
	var logs *LogBuffer
	if isValidating {
		logs = trackLogs()
	}
	report.Info("daemon.start", report.Data{
		"day":          day.Format("2006-01-02"),
		"project_id":   projectId,
		"pubsub_topic": topicName,
	})

	// download & parse CSV data
	snapshots, err := download(day)
	if err != nil {
		report.Action("download.fail", report.Data{"error": err.Error()})
		return err
	}
	report.Info("download.ok", report.Data{"count": len(snapshots)})

	// validate log stream on shutdown if required
	report.Drain()
	err = nil
	if isValidating {
		expect := map[string]int{
			"download.ok": VALIDATE_IS_PRESENT,
		}
		err = validateLogStream(logs, expect)
	}
	return err
}
