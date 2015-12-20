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
	demoMode := true

	refSnapC := make(chan *Snapshot, 10)
	for url := range discoverStationUrls() {
		_, measures := requestStationDetail(url)
		for _, m := range measures {
			refSnapC <- &m
		}
		if demoMode {
			break
		}
	}

	updateSnapC := make(chan *SnapshotUpdate, 10)
	go func() {
		tick := time.Tick(time.Minute * 15)
		for {
			requestLatestReadings(updateSnapC)
			<-tick
		}
	}()

	//pubSnapC := applyUpdatesToReferenceSnapshots(refSnapC, updateSnapC)
}
