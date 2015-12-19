package main

import (
	"time"

	"github.com/robtuley/report"
)

type GaugeSnapshot struct {
	WiskiId   string
	Name      string
	RiverName string
	Lat       float32
	Lg        float32
}

func main() {
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "ea_monitor"})
	report.RuntimeStatsEvery(30 * time.Second)

	gaugeSnapshotC := make(chan GaugeSnapshot)

	reTryWiskiIdC := make(chan string, 5000)
	for id := range discoverWiskiIDs() {
		requestStationDetail(id, reTryWiskiIdC, gaugeSnapshotC)
	}
	failedWiskiIdC := make(chan string, 5000)
	for id := range reTryWiskiIdC {
		requestStationDetail(id, failedWiskiIdC, gaugeSnapshotC)
	}

}
