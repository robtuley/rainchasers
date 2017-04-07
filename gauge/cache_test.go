package gauge

import (
	"reflect"
	"testing"
	"time"
)

func TestConcatDeDuplicatesAndOrdersByTime(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-01-01T10:30:00Z")
	r3 := Reading{
		DateTime: timestamp,
		Value:    3.21,
	}
	r2 := Reading{
		DateTime: timestamp.Add(time.Second),
		Value:    2.13,
	}
	r1 := Reading{
		DateTime: timestamp.Add(time.Second * 10),
		Value:    1.23,
	}

	var result []Reading

	result = concat([]Reading{r1, r3}, []Reading{r2, r1})
	if !reflect.DeepEqual(result, []Reading{r1, r2, r3}) {
		t.Error("concat() unexpected result", result)
	}

	result = concat([]Reading{r1, r2, r3}, []Reading{r2, r3, r1})
	if !reflect.DeepEqual(result, []Reading{r1, r2, r3}) {
		t.Error("concat() unexpected result", result)
	}

	result = concat([]Reading{}, []Reading{r3, r1})
	if !reflect.DeepEqual(result, []Reading{r1, r3}) {
		t.Error("concat() unexpected result", result)
	}

	result = concat([]Reading{r3, r2, r1}, []Reading{})
	if !reflect.DeepEqual(result, []Reading{r1, r2, r3}) {
		t.Error("concat() unexpected result", result)
	}
}

func TestRemoveOlderThan(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-01-01T10:30:00Z")
	r1 := Reading{
		DateTime: timestamp,
		Value:    1.0,
	}
	r2 := Reading{
		DateTime: timestamp.Add(time.Second),
		Value:    2.0,
	}
	r3 := Reading{
		DateTime: timestamp.Add(10 * time.Second),
		Value:    3.0,
	}
	r4 := Reading{
		DateTime: timestamp.Add(-1 * time.Second),
		Value:    4.0,
	}

	var result []Reading

	result = []Reading{r1, r2, r3, r4}
	removeOlderThan(timestamp, &result)
	if !reflect.DeepEqual(result, []Reading{r2, r3}) {
		t.Error("removeOlderThan unexpected result #1", result)
	}

	result = []Reading{r3, r1, r2}
	removeOlderThan(timestamp, &result)
	if !reflect.DeepEqual(result, []Reading{r3, r2}) {
		t.Error("removeOlderThan unexpected result #2", result)
	}

	result = []Reading{r2, r3}
	removeOlderThan(timestamp, &result)
	if !reflect.DeepEqual(result, []Reading{r2, r3}) {
		t.Error("removeOlderThan unexpected result #3", result)
	}

	result = []Reading{r4}
	removeOlderThan(timestamp, &result)
	if !reflect.DeepEqual(result, []Reading{}) {
		t.Error("removeOlderThan unexpected result #4", result)
	}
}
