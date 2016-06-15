package report

import (
	"time"
)

// Tock reports timer telemetry data recording the time since the Tick.
//
//     defer report.Tock(report.Tick(), "mongo.query", report.Data{"q":query})
//
func Tock(start time.Time, event string, payload Data) {
	payload["type"] = "timer"
	payload["event"] = event
	payload["ns"] = time.Now().Sub(start).Nanoseconds()

	publishRawEvent(payload)
}

// Tick starts a timer event with a value for the later call to Tock
func Tick() time.Time {
	return time.Now()
}
