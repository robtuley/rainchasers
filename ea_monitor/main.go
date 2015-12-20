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

	//pubSnapC := applyUpdatesToReferenceSnapshots(refSnapC, updateSnapC)
}
