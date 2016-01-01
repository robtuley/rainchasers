package report

import (
	"time"
)

// Report timer telemetry data: setup to record a function duration length.
//
//     defer report.Tock(report.Tick(), "mongo.query", report.Data{"q":query})
//
func Tock(start time.Time, event string, payload Data) {
	payload["type"] = "timer"
	payload["event"] = event
	payload["ns"] = time.Now().Sub(start).Nanoseconds()

	publishRawEvent(payload)
}

func Tick() time.Time {
	return time.Now()
}
