package main

import (
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
)

// River is the river as written to firestore
type River struct {
	Section  Section   `firestore:"section"`
	Level    *Level    `firestore:"level,omitempty"`
	Measures []Measure `firestore:"measures"`
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
	Station       gauge.Station   `firestore:"station"`
	Calibration   Calibration     `firestore:"calibration"`
	Readings      []gauge.Reading `firestore:"readings"`
	ProcessedTime time.Time       `firestore:"processed_time"`
}
