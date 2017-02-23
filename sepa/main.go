package main

import (
	"os"
	"strconv"
	"time"

	"github.com/robtuley/report"
)

// Responds to environment variables:
//   UPDATE_EVERY_X_SECONDS (default 15*60)
//   UPDATE_COUNT_BEFORE_SHUTDOWN (default 100)
//   PROJECT_ID (no default, blank for validation mode)
//   LATEST_PUBSUB_TOPIC (no default, blank for validation mode)
//   HISTORY_PUBSUB_TOPIC (no default, blank for validation mode)
//
func main() {
	if err := run(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run() error {
	// setup telemetry and logging
	defer report.Drain()
	report.StdOut()
	report.Global(report.Data{"service": "sepa", "daemon": time.Now().Format("v2006-01-02-15-04-05")})
	report.RuntimeStatsEvery(30 * time.Second)

	// parse env vars
	updatePeriodSeconds, err := strconv.Atoi(os.Getenv("UPDATE_EVERY_X_SECONDS"))
	if err != nil {
		updatePeriodSeconds = 15 * 60
	}
	updateCountOnShutdown, err := strconv.Atoi(os.Getenv("UPDATE_COUNT_BEFORE_SHUTDOWN"))
	if err != nil {
		updateCountOnShutdown = 100
	}
	projectId := os.Getenv("PROJECT_ID")
	latestTopicName := os.Getenv("LATEST_PUBSUB_TOPIC")
	historyTopicName := os.Getenv("HISTORY_PUBSUB_TOPIC")

	_ = updateCountOnShutdown
	_ = updatePeriodSeconds
	_ = projectId
	_ = latestTopicName
	_ = historyTopicName

	// discover SEPA gauging stations
	_, err = discoverStations()
	if err != nil {
		return err
	}

	//

	//
	// SEPA_HYDROLOGY_OFFICE,STATION_NAME,LOCATION_CODE,NATIONAL_GRID_REFERENCE,CATCHMENT_NAME,RIVER_NAME,GAUGE_DATUM,CATCHMENT_AREA,START_DATE,END_DATE,SYSTEM_ID,LOWEST_VALUE,LOW,MAX_VALUE,HIGH,MAX_DISPLAY,MEAN,UNITS,WEB_MESSAGE,NRFA_LINK
	// Perth,Perth,10048,NO1160525332,---,Tay,2.08,4991.0,August 1991,2017-02-20 12:45:00,58156010,0.0,0.168,4.928,3.493,4.928m @ 17/01/1993 19:30:00,0.894,m,,http://www.ceh.ac.uk/data/nrfa/data/station.html?15042
	//
	// 2. Calculate tick rate (with min) and spawn individual gauge download CSVs
	// 3. Download individual CSV, latest value to latest reading topic, previous to history
	// 4. Close and restart

	return nil
}
