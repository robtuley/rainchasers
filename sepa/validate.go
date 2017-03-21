package main

import (
	"errors"
	"github.com/rainchasers/report"
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

func (b LogBuffer) FirstActionableError() error {
	var err error
	b.Mutex.Lock()
	for name, events := range b.Map {
		for _, e := range events {
			if e["type"] == "action" {
				err = errors.New("Actionable error: " + name)
			}
		}
	}
	b.Mutex.Unlock()
	return err
}

func (b LogBuffer) Count() map[string]int {
	count := make(map[string]int)
	b.Mutex.Lock()
	for name, events := range b.Map {
		count[name] = len(events)
	}
	b.Mutex.Unlock()
	return count
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
	err := buffer.FirstActionableError()
	if err != nil {
		return err
	}

	count := buffer.Count()
	for k, expect := range expectCounts {
		val, exists := count[k]
		if !exists {
			return errors.New(k + " expected but not present")
		}
		if expect != VALIDATE_IS_PRESENT && expect != val {
			return errors.New(k + " unexpected count")
		}
	}
	return nil
}
