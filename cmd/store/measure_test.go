package main

import (
	"reflect"
	"testing"
	"time"

	"github.com/robtuley/rainchasers/internal/gauge"
)

func TestMergeDeDuplicatesAndOrdersByTime(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-01-01T10:30:00Z")
	r3 := gauge.Reading{
		EventTime: timestamp,
		Value:     3.21,
	}
	r2 := gauge.Reading{
		EventTime: timestamp.Add(time.Second),
		Value:     2.13,
	}
	r1 := gauge.Reading{
		EventTime: timestamp.Add(time.Second * 10),
		Value:     1.23,
	}

	var result []gauge.Reading

	result = merge([]gauge.Reading{r1, r3}, []gauge.Reading{r2, r1})
	if !reflect.DeepEqual(result, []gauge.Reading{r1, r2, r3}) {
		t.Error("merge() unexpected result", result)
	}

	result = merge([]gauge.Reading{r1, r2, r3}, []gauge.Reading{r2, r3, r1})
	if !reflect.DeepEqual(result, []gauge.Reading{r1, r2, r3}) {
		t.Error("merge() unexpected result", result)
	}

	result = merge([]gauge.Reading{}, []gauge.Reading{r3, r1})
	if !reflect.DeepEqual(result, []gauge.Reading{r1, r3}) {
		t.Error("merge() unexpected result", result)
	}

	result = merge([]gauge.Reading{r3, r2, r1}, []gauge.Reading{})
	if !reflect.DeepEqual(result, []gauge.Reading{r1, r2, r3}) {
		t.Error("merge() unexpected result", result)
	}

	result = merge(nil, []gauge.Reading{r2, r1})
	if !reflect.DeepEqual(result, []gauge.Reading{r1, r2}) {
		t.Error("merge() unexpected result", result)
	}
}

func TestRemoveOlderThan(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-01-01T10:30:00Z")
	r1 := gauge.Reading{
		EventTime: timestamp,
		Value:     1.0,
	}
	r2 := gauge.Reading{
		EventTime: timestamp.Add(time.Second),
		Value:     2.0,
	}
	r3 := gauge.Reading{
		EventTime: timestamp.Add(10 * time.Second),
		Value:     3.0,
	}
	r4 := gauge.Reading{
		EventTime: timestamp.Add(-1 * time.Second),
		Value:     4.0,
	}

	var result []gauge.Reading

	result = []gauge.Reading{r1, r2, r3, r4}
	removeOlderThan(timestamp, &result)
	if !reflect.DeepEqual(result, []gauge.Reading{r2, r3}) {
		t.Error("removeOlderThan unexpected result #1", result)
	}

	result = []gauge.Reading{r3, r1, r2}
	removeOlderThan(timestamp, &result)
	if !reflect.DeepEqual(result, []gauge.Reading{r3, r2}) {
		t.Error("removeOlderThan unexpected result #2", result)
	}

	result = []gauge.Reading{r2, r3}
	removeOlderThan(timestamp, &result)
	if !reflect.DeepEqual(result, []gauge.Reading{r2, r3}) {
		t.Error("removeOlderThan unexpected result #3", result)
	}

	result = []gauge.Reading{r4}
	removeOlderThan(timestamp, &result)
	if !reflect.DeepEqual(result, []gauge.Reading{}) {
		t.Error("removeOlderThan unexpected result #4", result)
	}
}
