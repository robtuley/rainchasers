package main

import (
	"time"

	"github.com/robtuley/report"
)

func main() {
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "ea_monitor"})
	report.RuntimeStatsEvery(30 * time.Second)
	isDemo := true

	refSnapC := make(chan Snapshot, 10)
	updateSnapC := make(chan SnapshotUpdate, 10)
	pubSnapC := applyUpdatesToRefSnaps(refSnapC, updateSnapC)

	// blackhole pubSnapC (todo: send to Google pub/sub)
	go func() {
		for s := range pubSnapC {
			report.Info("published", report.Data{"snapshot": s})
		}
	}()

	// retrieve list of all stations & latest readings
	for url := range discoverStationUrls() {
		_, measures := requestStationDetail(url)
		for _, m := range measures {
			refSnapC <- m
		}
		if isDemo {
			break
		}
	}

	// poll for latest readings
	tick := time.Tick(time.Minute * 15)
	for {
		requestLatestReadings(updateSnapC)
		<-tick
	}
}
