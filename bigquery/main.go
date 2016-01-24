package main

import (
	"os"
	"strconv"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/robtuley/report"
)

type actionableEvent struct {
	EventName string
	Message   string
}

// Responds to environment variables:
//   GCLOUD_PROJECT_ID (no default)
//   GCLOUD_PUBSUB_TOPIC (no default)
//   GCLOUD_BUCKET_NAME (no default)
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
	bucketName := os.Getenv("GCLOUD_BUCKET_NAME")
	batchSize, err := strconv.Atoi(os.Getenv("SNAPSHOT_BATCH_SIZE"))
	if err != nil {
		batchSize = 1000
	}

	report.Info("daemon.start", report.Data{
		"project_id":   projectId,
		"pubsub_topic": topicName,
		"bucket_name":  bucketName,
		"batch_size":   batchSize,
	})

	// setup actionable events channel
	actionC := make(chan actionableEvent)

	// consume snapshots from pubsub
	ctx, err := gauge.NewPubSubContext(projectId, topicName)
	if err != nil {
		report.Action("connect.error", report.Data{"error": err.Error()})
		return
	}
	snapC, snapErrC, err := gauge.Subscribe(ctx, "snap-to-bigquery")
	if err != nil {
		report.Action("consume.error", report.Data{"error": err.Error()})
		return
	}
	go func() {
		for err := range snapErrC {
			actionC <- actionableEvent{"consume.error", err.Error()}
		}
	}()

	// preliminary in-memory de-dup
	dedupC := make(chan gauge.Snapshot, 10)
	go func() {
		cache := newDeDupeCache(10000)
		for s := range snapC {
			id := s.InsertId()
			if !cache.Exists(id) {
				dedupC <- s
			}
			cache.Set(id)
		}
	}()

	// buffer in-memory, flush to long-term CSV file storage
	csvC, csvErrC, err := csvEncodeAndWrite(projectId, bucketName, batchSize, dedupC)
	if err != nil {
		report.Action("csv.error", report.Data{"error": err.Error()})
		return
	}
	go func() {
		for err := range csvErrC {
			actionC <- actionableEvent{"csv.error", err.Error()}
		}
	}()

	// todo: load CSV file into bigquery table then perform
	// final dedup and load queries
	bqErrC, err := loadCSVIntoBigQuery(projectId, "rainchasers", "gauge_reading", csvC)
	if err != nil {
		report.Action("bigquery.error", report.Data{"error": err.Error()})
		return
	}
	go func() {
		for err := range bqErrC {
			actionC <- actionableEvent{"bigquery.error", err.Error()}
		}
	}()

	// log any actionable events.
	for e := range actionC {
		report.Action(e.EventName, report.Data{"error": e.Message})
	}

	report.Info("daemon.stop", report.Data{})
}
