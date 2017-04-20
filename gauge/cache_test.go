package gauge

import (
	"context"
	"reflect"
	"testing"
	"time"
)

func TestConcatDeDuplicatesAndOrdersByTime(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-01-01T10:30:00Z")
	r3 := Reading{
		EventTime: timestamp,
		Value:     3.21,
	}
	r2 := Reading{
		EventTime: timestamp.Add(time.Second),
		Value:     2.13,
	}
	r1 := Reading{
		EventTime: timestamp.Add(time.Second * 10),
		Value:     1.23,
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
		EventTime: timestamp,
		Value:     1.0,
	}
	r2 := Reading{
		EventTime: timestamp.Add(time.Second),
		Value:     2.0,
	}
	r3 := Reading{
		EventTime: timestamp.Add(10 * time.Second),
		Value:     3.0,
	}
	r4 := Reading{
		EventTime: timestamp.Add(-1 * time.Second),
		Value:     4.0,
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

func TestCacheAddAndGet(t *testing.T) {
	timestamp := time.Now()

	r1 := Reading{
		EventTime: timestamp,
		Value:     3.21,
	}
	r2 := Reading{
		EventTime: timestamp.Add(-1 * time.Second),
		Value:     2.13,
	}
	r3 := Reading{
		EventTime: timestamp.Add(-2 * time.Second),
		Value:     1.23,
	}

	stationA := Station{
		DataURL: "http://example.com/A",
	}
	stationB := Station{
		DataURL: "http://example.com/B",
	}

	stationAsnap1 := Snapshot{
		Station:  stationA,
		Readings: []Reading{r1, r3},
	}

	stationAsnap2 := Snapshot{
		Station:  stationA,
		Readings: []Reading{r1, r2, r3},
	}

	stationBsnap1 := Snapshot{
		Station:  stationB,
		Readings: []Reading{r1, r2},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cache := NewCache(ctx, time.Hour)

	stat := cache.Stats()
	if stat.StationCount != 0 || stat.MaxReadingCount != 0 || stat.MinReadingCount != 0 {
		t.Error("Empty cache stat count mismatch", stat)
	}
	if stat.OldestReading.Seconds() != 0 {
		t.Error("Empty cache oldest time mismatch", stat)
	}

	cache.Add(&stationAsnap1)
	cache.Add(&stationBsnap1)
	cache.Add(&stationAsnap2)

	stationAresult, exists := cache.Get(stationA.UUID())
	if !exists {
		t.Error("Station A uncached", stationA.UUID())
	} else {
		if !reflect.DeepEqual(stationAresult.Station, stationA) {
			t.Error("Station mismatch for A", stationAresult.Station)
		}
		if !reflect.DeepEqual(stationAresult.Readings, concat(stationAsnap1.Readings, stationAsnap2.Readings)) {
			t.Error("Readings mismatch for A", stationAresult.Readings)
		}
		if !stationAresult.ProcessingTime.After(timestamp) {
			t.Error("Incorrect Modified at for A", stationAresult.ProcessingTime)
		}
	}

	stationBresult, exists := cache.Get(stationB.UUID())
	if !exists {
		t.Error("Station B uncached", stationB.UUID())
	} else {
		if !reflect.DeepEqual(stationBresult.Station, stationB) {
			t.Error("Station mismatch for B", stationBresult.Station)
		}
		if !reflect.DeepEqual(stationBresult.Readings, stationBsnap1.Readings) {
			t.Error("Readings mismatch for B", stationBresult.Readings)
		}
		if !stationBresult.ProcessingTime.After(timestamp) {
			t.Error("Incorrect Modified at for B", stationBresult.ProcessingTime)
		}
	}

	stat = cache.Stats()
	if stat.StationCount != 2 {
		t.Error("Cache stat station count mismatch", stat)
	}
	if stat.AllReadingCount != 5 {
		t.Error("Cache stat all reading count mismatch", stat)
	}
	if stat.MaxReadingCount != 3 {
		t.Error("Cache stat max reading count mismatch", stat)
	}
	if stat.MinReadingCount != 2 {
		t.Error("Cache stat max reading count mismatch", stat)
	}
	if stat.OldestReading.Seconds() < 2 {
		t.Error("Empty cache oldest time mismatch", stat)
	}
}
