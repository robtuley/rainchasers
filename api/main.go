package main

import (
	"github.com/rainchasers/report"
	"os"
	"time"
)

// Responds to environment variables:
//   BOOTSTRAP_URL (no default)
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
	report.Info("daemon.start", report.Data{
		"bootstrap_url": bootstrapURL,
		"project_id":    projectId,
		"pubsub_topic":  topicName,
	})

	return nil
}
