package main

import (
	"time"

	"github.com/robtuley/report"
)

const (
	API_REQUESTS_PER_SECOND = 1
	API_USER_AGENT          = "Rainchasers v0.1"
)

var apiRequestThrottle <-chan time.Time

func main() {
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "ea_monitor"})
	report.RuntimeStatsEvery(30 * time.Second)

	apiRequestThrottle = time.Tick(time.Second / API_REQUESTS_PER_SECOND)
	n := 0
	for range discoverWiskiIDs() {
		n = n + 1
	}
	report.Info("discover.found", report.Data{"count": n})
}
