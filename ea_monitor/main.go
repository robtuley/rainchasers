package main

import (
	"math"
	"os"
	"strconv"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"github.com/robtuley/report"
)

// Responds to 3 environment variables:
//   EA_MONITOR_DISCOVER_STATIONS_LIMIT (default no limit)
//   EA_MONITOR_UPDATE_EVERY_X_SECONDS (default 15*60)
//   EA_MONITOR_UPDATE_COUNT_BEFORE_SHUTDOWN (default 100)
//
func main() {

	// setup telemetry and logging
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "eam", "daemon": time.Now().Unix()})
	report.RuntimeStatsEvery(30 * time.Second)

	// parse env vars
	stationLimit, err := strconv.Atoi(os.Getenv("EA_MONITOR_DISCOVER_STATIONS_LIMIT"))
	if err != nil {
		stationLimit = math.MaxInt32
	}
	updatePeriodSeconds, err := strconv.Atoi(os.Getenv("EA_MONITOR_UPDATE_EVERY_X_SECONDS"))
	if err != nil {
		updatePeriodSeconds = 15 * 60
	}
	updateCountOnShutdown, err := strconv.Atoi(os.Getenv("EA_MONITOR_UPDATE_COUNT_BEFORE_SHUTDOWN"))
	if err != nil {
		updateCountOnShutdown = 100
	}
	report.Info("daemon.start", report.Data{
		"station_limit":            stationLimit,
		"update_period":            updatePeriodSeconds,
		"update_count_on_shutdown": updateCountOnShutdown,
	})

	// init in-bounds channels & publisher
	refSnapC := make(chan gauge.Snapshot, 10)
	updateSnapC := make(chan gauge.SnapshotUpdate, 10)
	pubSnapC := applyUpdatesToRefSnaps(refSnapC, updateSnapC)

	// blackhole pubSnapC (todo: send to Google pub/sub)
	go func() {
		for s := range pubSnapC {
			report.Info("gpubsub.snapshot", report.Data{"snapshot": s})
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
}
