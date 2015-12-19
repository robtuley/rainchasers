package main

import (
	"time"

	"github.com/robtuley/report"
)

type Snapshot struct {
	Url        string
	StationUrl string
	Name       string
	RiverName  string
	Lat        float32
	Lg         float32
	Type       string
	Unit       string
	DateTime   time.Time
	Value      float32
}

func main() {
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "ea_monitor"})
	report.RuntimeStatsEvery(30 * time.Second)

	snapshotC := make(chan *Snapshot, 10)

	for url := range discoverStationUrls() {
		_, measures := requestStationDetail(url)
		for _, m := range measures {
			snapshotC <- &m
		}
	}

}
