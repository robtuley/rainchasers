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
//   PROJECT_ID (no default)
//   PUBSUB_TOPIC (no default)
//
func main() {

	// setup telemetry and logging
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "eam", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
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
	projectId := os.Getenv("PROJECT_ID")
	topicName := os.Getenv("PUBSUB_TOPIC")
	report.Info("daemon.start", report.Data{
		"station_limit":            stationLimit,
		"update_period":            updatePeriodSeconds,
		"update_count_on_shutdown": updateCountOnShutdown,
		"project_id":               projectId,
		"pubsub_topic":             topicName,
	})

	// init in-bounds channels & publisher
	refSnapC := make(chan gauge.Snapshot, 10)
	updateLatestC := make(chan gauge.SnapshotUpdate, 10)
	updateHistoryC := make(chan gauge.SnapshotUpdate, 10)

	latestSnapC, historySnapC := applyUpdatesToRefSnaps(refSnapC, updateLatestC, updateHistoryC)

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
			case s, is_ok := <-latestSnapC:
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
			case s, is_ok := <-historySnapC:
				if !is_ok {
					break
				}
				report.Info("history.snapshot", report.Data{
					"snapshot": s,
				})
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

	// start job to download historical data
	historyErrC := downloadHistoricalDataForDaysAgo(2, updateHistoryC)
	go func() {
		for err := range historyErrC {
			report.Action("history.error", report.Data{"error": err.Error()})
		}
	}()

	// poll for latest readings
	tick := time.Tick(time.Second * time.Duration(updatePeriodSeconds))
	n := 0
	for {
		requestLatestReadings(updateLatestC)
		n = n + 1
		if n == updateCountOnShutdown {
			break
		}
		<-tick
	}

	close(updateLatestC)
	close(updateHistoryC)
	report.Info("daemon.stop", report.Data{})
}
