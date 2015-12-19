package main

import (
	"time"

	"github.com/robtuley/report"
)

type GaugeSnapshot struct {
	Url       string
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

	reTryUrlC := make(chan string, 10)
	for url := range discoverUrls() {
		requestStationDetail(url, reTryUrlC, gaugeSnapshotC)
	}
	for url := range reTryUrlC() {
		requestStationDetail(url, reTryUrlC, gaugeSnapshotC)
	}
}
