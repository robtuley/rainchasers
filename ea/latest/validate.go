package main

import (
	"errors"

	"github.com/robtuley/report"
)

const VALIDATE_IS_PRESENT = iota * -1

var logEventMap map[string][]report.Data

func init() {
	logEventMap = make(map[string][]report.Data)
}

func trackLogs() {
	report.Observe(func(d report.Data) {
		evtName := d["event"].(string)
		logEventMap[evtName] = append(logEventMap[evtName], d)
	})
}

func validateLogStream(expectEventCounts map[string]int) error {
	n := 0
	count := make(map[string]int)

	for name, evts := range logEventMap {
		count[name] = len(evts)

		// detect and fail on any actionable errors
		for _, e := range evts {
			if e["type"] == "action" {
				return errors.New("Actionable error: " + name)
			}
		}
	}
	report.Info("daemon.validation", report.Data{"buffer_size": n, "count": count})

	// make assertions on event counts
	for k, expect := range expectEventCounts {
		val, exists := count[k]
		if !exists {
			return errors.New(k + " not present")
		}
		if expect != VALIDATE_IS_PRESENT && expect != val {
			return errors.New(k + " unexpected count")
		}
	}

	return nil
}
