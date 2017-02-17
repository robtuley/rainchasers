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
//   LATEST_PUBSUB_TOPIC (no default)
//   HISTORY_PUBSUB_TOPIC (no default)
//   BUCKET_NAME (no default)
//   BIGQUERY_DATASET (no default)
//   BIGQUERY_TABLE (no default)
//   MIN_BATCH_SIZE (default 1000)
//   MIN_BATCH_SECONDS (default 20*60)
//   ERROR_COUNT_ON_EXIT (default 10)
//
func main() {

	// setup telemetry and logging
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "bigq", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	report.RuntimeStatsEvery(30 * time.Second)

	// parse env vars
	projectId := os.Getenv("PROJECT_ID")
	latestTopicName := os.Getenv("LATEST_PUBSUB_TOPIC")
	historyTopicName := os.Getenv("HISTORY_PUBSUB_TOPIC")
	bucketName := os.Getenv("BUCKET_NAME")
	datasetId := os.Getenv("BIGQUERY_DATASET")
	tableId := os.Getenv("BIGQUERY_TABLE")
	batchMinSize, err := strconv.Atoi(os.Getenv("MIN_BATCH_SIZE"))
	if err != nil {
		batchMinSize = 1000
	}
	batchMinSeconds, err := strconv.Atoi(os.Getenv("MIN_BATCH_SECONDS"))
	if err != nil {
		batchMinSeconds = 20 * 60
	}
	errCountOnExit, err := strconv.Atoi(os.Getenv("ERROR_COUNT_ON_EXIT"))
	if err != nil {
		errCountOnExit = 10
	}

	report.Info("daemon.start", report.Data{
		"project_id":           projectId,
		"latest_pubsub_topic":  latestTopicName,
		"history_pubsub_topic": historyTopicName,
		"bucket_name":          bucketName,
		"min_batch_size":       batchMinSize,
		"min_batch_seconds":    batchMinSeconds,
		"dataset":              datasetId,
		"table":                tableId,
	})

	// setup actionable events channel
	actionC := make(chan actionableEvent)

	// consume snapshots from pubsub
	snapC, snapErrC, err := consumeTopics(projectId, latestTopicName, historyTopicName)
	if err != nil {
		report.Action("error.subscribe", report.Data{"error": err.Error()})
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
	csvC, csvErrC := csvEncodeAndWrite(projectId, bucketName, batchMinSize, batchMinSeconds, dedupC)
	go func() {
		for err := range csvErrC {
			actionC <- actionableEvent{"error.csv", err.Error()}
		}
	}()

	// load CSV file into bigquery table
	batchStatusC := loadCSVIntoBigQuery(projectId, datasetId, tableId+"_recent_with_dups", csvC)
	go func() {
		for s := range batchStatusC {
			report.Info("job.status", report.Data{
				"file": s.File,
				"jobs": s.Jobs,
			})
			if s.Error != nil {
				actionC <- actionableEvent{"error.bigquery", s.Error.Error()}
			}
		}
	}()

	// log any actionable events.
	nErr := 0
	for e := range actionC {
		nErr += 1
		report.Action(e.EventName, report.Data{"error": e.Message})
		if nErr > errCountOnExit {
			break
		}
	}

	report.Info("daemon.stop", report.Data{})
}
