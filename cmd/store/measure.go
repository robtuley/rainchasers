package main

import (
	"sort"
	"strconv"
	"time"

	"github.com/robtuley/rainchasers/internal/gauge"
	"github.com/robtuley/rainchasers/internal/river"
)

type readingSorter []gauge.Reading

func (rs readingSorter) Len() int {
	return len(rs)
}

func (rs readingSorter) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs readingSorter) Less(i, j int) bool {
	return rs[i].EventTime.After(rs[j].EventTime)
}

// Record is the river as written to firestore
type Record struct {
	Section  river.Section `firestore:"section"`
	Level    Level         `firestore:"level"`
	Measures []Measure     `firestore:"measures"`
}

// Level is the calculated river current level state
type Level struct {
	EventTime     time.Time `firestore:"event_time"`     // time when this was true
	ProcessedTime time.Time `firestore:"processed_time"` // time when this was processed
	Label         string    `firestore:"label"`          // e.g. "high"
	Reason        string    `firestore:"reason"`         // e.g. "0.98 at Hafod Wydr gauge"
}

// Measure is a relevant river measurement time series
type Measure struct {
	Station       gauge.Station     `firestore:"station"`
	Calibration   river.Calibration `firestore:"calibration"`
	Readings      []gauge.Reading   `firestore:"readings"`
	ProcessedTime time.Time         `firestore:"processed_time"`
}

// LatestLevel returns the latest level calibration if available
func (m Measure) LatestLevel() Level {
	if len(m.Readings) == 0 {
		return Level{
			Label:  river.Unknown.String(),
			Reason: "No readings currently available from " + m.Station.Name,
		}
	}

	// readings are always sorted by most recent first by convention
	r := m.Readings[0]
	v := strconv.FormatFloat(float64(r.Value), 'f', 2, 32)
	return Level{
		EventTime:     r.EventTime,
		ProcessedTime: m.ProcessedTime,
		Label:         m.Calibration.LevelAt(r.Value).String(),
		Reason:        v + " at " + m.Station.Name,
	}
}

func merge(a []gauge.Reading, b []gauge.Reading) []gauge.Reading {
	concat := append(a, b...)
	removeDuplicates(&concat)
	sort.Sort(readingSorter(concat))
	return concat
}

func removeDuplicates(xs *[]gauge.Reading) {
	found := make(map[time.Time]bool)
	j := 0
	for i, x := range *xs {
		if !found[x.EventTime] {
			found[x.EventTime] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

func removeOlderThan(keepSince time.Time, xs *[]gauge.Reading) {
	j := 0
	for i, x := range *xs {
		if keepSince.Before(x.EventTime) {
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}
