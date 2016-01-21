package main

import (
	"os"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/robtuley/report"
)

// Responds to environment variables:
//   GCLOUD_PROJECT_ID (no default)
//   GCLOUD_PUBSUB_TOPIC (no default)
//
func main() {

	// setup telemetry and logging
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "bigq", "daemon": time.Now().Unix()})
	report.RuntimeStatsEvery(30 * time.Second)

	// parse env vars
	projectId := os.Getenv("GCLOUD_PROJECT_ID")
	topicName := os.Getenv("GCLOUD_PUBSUB_TOPIC")
	report.Info("daemon.start", report.Data{
		"project_id":   projectId,
		"pubsub_topic": topicName,
	})

	// consume snapshots from pubsub
	ctx, err := gauge.NewContext(projectId, topicName)
	if err != nil {
		report.Action("gpubsub.connect.error", report.Data{"error": err.Error()})
		return
	}
	snapC, errC, err := gauge.Consume(ctx, "snap-to-bigquery")
	if err != nil {
		report.Action("gpubsub.consume.error", report.Data{"error": err.Error()})
		return
	}
	go func() {
		for err := range errC {
			report.Action("gpubsub.consume.error", report.Data{"error": err.Error()})
		}
	}()

	// consume, de-dup & encode to csv, writing every
	// hour to cloud storage
	dedup := newDeDupeCache(10000)
	for s := range snapC {
		id := s.InsertId()
		if !dedup.Exists(id) {
			report.Info("gpubsub.consume.ok", report.Data{"snapshot": s})
		}
		dedup.Set(id)
	}

	report.Info("daemon.stop", report.Data{})
}
