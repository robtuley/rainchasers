package main

import (
	"errors"
	"github.com/robtuley/report"
	"sync"
)

const VALIDATE_IS_PRESENT = iota * -1

type LogBuffer struct {
	Mutex *sync.Mutex
	Map   map[string][]report.Data
}

func (b LogBuffer) Add(d report.Data) {
	evtName := d["event"].(string)
	b.Mutex.Lock()
	b.Map[evtName] = append(b.Map[evtName], d)
	b.Mutex.Unlock()
}

func (b LogBuffer) Extract() map[string][]report.Data {
	b.Mutex.Lock()
	logs := b.Map
	b.Mutex.Unlock()
	return logs
}

func trackLogs() *LogBuffer {
	buffer := LogBuffer{
		Mutex: &sync.Mutex{},
		Map:   make(map[string][]report.Data),
	}
	report.Observe(buffer.Add)
	return &buffer
}

func validateLogStream(buffer *LogBuffer, expectCounts map[string]int) error {
	n := 0
	count := make(map[string]int)

	for name, evts := range buffer.Extract() {
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
	for k, expect := range expectCounts {
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
