package main

import (
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
	"sort"
	"time"

	"github.com/rainchasers/content/internal/gauge"
	"github.com/rainchasers/content/internal/river"
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
	Level    *Level        `firestore:"level,omitempty"`
	Measures []Measure     `firestore:"measures"`
}

// Level is the calculated river current state
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

// Checksum provides a checksum of current state
func (m Measure) Checksum() string {
	b := sha1.New()
	gob.NewEncoder(b).Encode(m)
	return hex.EncodeToString(b.Sum(nil))
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
