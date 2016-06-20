package main

import (
	"os"
	"strconv"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
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
	report.Global(report.Data{"service": "eam", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
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

	// init in-bounds channels & publisher
	refSnapC := make(chan gauge.Snapshot, 10)
	updateLatestC := make(chan gauge.SnapshotUpdate, 10)
	updateHistoryC := make(chan gauge.SnapshotUpdate, 10)

	// combine updates (time & value) with reference snapshot data (everything else)
	latestSnapC, historySnapC := applyUpdatesToRefSnaps(refSnapC, updateLatestC, updateHistoryC)

	// publish snapshots to latest & history PubSub topic
	if isValidating {
		go logSnapshotsFromChannel("snapshot.latest", latestSnapC)
		go logSnapshotsFromChannel("snapshot.history", historySnapC)
	} else {
		err = publishSnapshotsFromChannels(projectId, latestTopicName, historyTopicName, latestSnapC, historySnapC)
		if err != nil {
			report.Action("pubsub.connect.error", report.Data{"error": err.Error()})
			return err
		}
	}

	// retrieve list of all stations & latest readings
	var stationC chan string
	if isValidating {
		stationC = sampleStationUrls()
	} else {
		stationC = discoverStationUrls()
	}
	for url := range stationC {
		measures, _ := requestStationDetail(url)
		for _, m := range measures {
			refSnapC <- m
		}
	}

	// start job to download historical data
	historyErrC := downloadHistoricalDataForDaysAgo(1, updateHistoryC)
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

	// validate log stream on shutdown if required
	err = nil
	if isValidating {
		expect := map[string]int{
			"update.response":  updateCountOnShutdown,
			"snapshot.history": VALIDATE_IS_PRESENT,
			"snapshot.latest":  VALIDATE_IS_PRESENT,
		}
		err = validateLogStream(validateC, expect)
	}
	return err
}
