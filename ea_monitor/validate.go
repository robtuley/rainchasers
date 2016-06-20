package main

import (
	"errors"

	"github.com/robtuley/report"
)

const (
	VALIDATE_IS_PRESENT = iota * -1
)

func bufferLogStream(bufferSize int) <-chan report.Data {
	logC := make(chan report.Data, bufferSize)
	log2Channel := func(d report.Data) {
		select {
		case logC <- d:
		default:
			// (a non-blocking publish)
		}
	}
	report.Observe(log2Channel)
	return logC
}

func validateLogStream(logC <-chan report.Data, expectEventCounts map[string]int) error {
	n := 0
	count := make(map[string]int)

Consume:
	for {
		select {
		case d := <-logC:
			n = n + 1
			evtName := d["event"].(string)

			// detect and fail on any actionable errors
			if d["type"] == "action" {
				return errors.New("Actionable error: " + evtName)
			}

			// count types of each event
			count[evtName] = count[evtName] + 1
		default:
			break Consume
		}
	}
	report.Info("daemon.validation", report.Data{"buffer_size": n, "count": count})

	// if buffer is full, we likely discarded logs
	if n == cap(logC) {
		return errors.New("Buffer capacity exceeded")
	}

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
