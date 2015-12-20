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

	refSnapC := make(chan *Snapshot, 10)
	for url := range discoverStationUrls() {
		_, measures := requestStationDetail(url)
		for _, m := range measures {
			refSnapC <- &m
		}
	}

	updateSnapC := make(chan *SnapshotUpdate, 10)
	ticker := time.NewTicker(time.Minute * 15)
	go func() {
		for range ticker.C {
			requestLatestReadings(updateSnapC)
		}
	}()

	//pubSnapC := applyUpdatesToReferenceSnapshots(refSnapC, updateSnapC)
}
