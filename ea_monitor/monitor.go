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

	for id := range discoverWiskiIDs() {
		_ = id
		// stationDetail(id)
	}
}
