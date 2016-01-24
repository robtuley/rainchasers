package main

import (
	"math"
	"os"
	"strconv"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/robtuley/report"
)

// Responds to environment variables:
//   DISCOVER_STATIONS_LIMIT (default no limit)
//   UPDATE_EVERY_X_SECONDS (default 15*60)
//   UPDATE_COUNT_BEFORE_SHUTDOWN (default 100)
//   GCLOUD_PROJECT_ID (no default)
//   GCLOUD_PUBSUB_TOPIC (no default)
//
func main() {

	// setup telemetry and logging
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "eam", "daemon": time.Now().Unix()})
	report.RuntimeStatsEvery(30 * time.Second)

	// parse env vars
	stationLimit, err := strconv.Atoi(os.Getenv("DISCOVER_STATIONS_LIMIT"))
	if err != nil {
		stationLimit = math.MaxInt32
	}
	updatePeriodSeconds, err := strconv.Atoi(os.Getenv("UPDATE_EVERY_X_SECONDS"))
	if err != nil {
		updatePeriodSeconds = 15 * 60
	}
	updateCountOnShutdown, err := strconv.Atoi(os.Getenv("UPDATE_COUNT_BEFORE_SHUTDOWN"))
	if err != nil {
		updateCountOnShutdown = 100
	}
	projectId := os.Getenv("GCLOUD_PROJECT_ID")
	topicName := os.Getenv("GCLOUD_PUBSUB_TOPIC")
	report.Info("daemon.start", report.Data{
		"station_limit":            stationLimit,
		"update_period":            updatePeriodSeconds,
		"update_count_on_shutdown": updateCountOnShutdown,
		"project_id":               projectId,
		"pubsub_topic":             topicName,
	})

	// init in-bounds channels & publisher
	refSnapC := make(chan gauge.Snapshot, 10)
	updateSnapC := make(chan gauge.SnapshotUpdate, 10)
	pubSnapC := applyUpdatesToRefSnaps(refSnapC, updateSnapC)

	// publish snapshots to PubSub topic
	ctx, err := gauge.NewPubSubContext(projectId, topicName)
	if err != nil {
		report.Action("pubsub.connect.error", report.Data{"error": err.Error()})
		return
	}
	go func() {
		tickC := time.Tick(time.Second * 10)
		n := 0
		for {
			select {
			case s, is_ok := <-pubSnapC:
				if !is_ok {
					break
				}
				err := gauge.Publish(ctx, s)
				n = n + 1
				if err != nil {
					report.Action("pubsub.publish.error", report.Data{
						"snapshot": s,
						"error":    err.Error(),
					})
				}
			case <-tickC:
				report.Info("pubsub.publish.ok", report.Data{"count": n})
				n = 0
			}
		}
	}()

	// retrieve list of all stations & latest readings
	for url := range discoverStationUrls(stationLimit) {
		measures, _ := requestStationDetail(url)
		for _, m := range measures {
			refSnapC <- m
		}
	}

	// poll for latest readings
	tick := time.Tick(time.Second * time.Duration(updatePeriodSeconds))
	n := 0
	for {
		requestLatestReadings(updateSnapC)
		n = n + 1
		if n == updateCountOnShutdown {
			break
		}
		<-tick
	}

	close(updateSnapC)
	close(pubSnapC)
	report.Info("daemon.stop", report.Data{})
}
