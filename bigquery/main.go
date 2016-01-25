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
//   PROJECT_ID (no default)
//   PUBSUB_TOPIC (no default)
//   BUCKET_NAME (no default)
//   BIGQUERY_DATASET (no default)
//   BIGQUERY_TABLE (no default)
//
func main() {

	// setup telemetry and logging
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "bigq", "daemon": time.Now().Unix()})
	report.RuntimeStatsEvery(30 * time.Second)

	// parse env vars
	projectId := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")
	bucketName := os.Getenv("BUCKET_NAME")
	datasetId := os.Getenv("BIGQUERY_DATASET")
	tableId := os.Getenv("BIGQUERY_TABLE")
	batchSize, err := strconv.Atoi(os.Getenv("SNAPSHOT_BATCH_SIZE"))
	if err != nil {
		batchSize = 1000
	}

	report.Info("daemon.start", report.Data{
		"project_id":   projectId,
		"pubsub_topic": topicName,
		"bucket_name":  bucketName,
		"batch_size":   batchSize,
		"dataset":      datasetId,
		"table":        tableId,
	})

	// setup actionable events channel
	actionC := make(chan actionableEvent)

	// consume snapshots from pubsub
	ctx, err := gauge.NewPubSubContext(projectId, topicName)
	if err != nil {
		report.Action("error.connect", report.Data{"error": err.Error()})
		return
	}
	snapC, snapErrC, err := gauge.Subscribe(ctx, "snap-to-bigquery")
	if err != nil {
		report.Action("error.consume", report.Data{"error": err.Error()})
		return
	}
	go func() {
		for err := range snapErrC {
			actionC <- actionableEvent{"error.consume", err.Error()}
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
		report.Action("error.csv", report.Data{"error": err.Error()})
		return
	}
	go func() {
		for err := range csvErrC {
			actionC <- actionableEvent{"error.csv", err.Error()}
		}
	}()

	// load CSV file into bigquery table
	batchStatusC, bqErrC, err := loadCSVIntoBigQuery(projectId, datasetId, tableId, csvC)
	if err != nil {
		report.Action("error.bigquery", report.Data{"error": err.Error()})
		return
	}
	go func() {
		for err := range bqErrC {
			actionC <- actionableEvent{"error.bigquery", err.Error()}
		}
	}()
	go func() {
		for s := range batchStatusC {
			report.Info("job.status", report.Data{
				"file": s.File,
				"jobs": s.Jobs,
			})
		}
	}()

	// log any actionable events.
	for e := range actionC {
		report.Action(e.EventName, report.Data{"error": e.Message})
	}

	report.Info("daemon.stop", report.Data{})
}
