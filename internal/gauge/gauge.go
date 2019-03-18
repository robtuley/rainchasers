package gauge

import (
	"time"
)

// Station is a specific location that can measure one or more metrics
type Station struct {
	DataURL   string  `json:"data_url"`
	AliasURL  string  `json:"alias_url"`
	HumanURL  string  `json:"human_url"`
	Name      string  `json:"name"`
	RiverName string  `json:"river"`
	Lat       float32 `json:"lat"`
	Lg        float32 `json:"lng"`
	Type      string  `json:"type"`
	Unit      string  `json:"unit"`
}

// Reading is a point-in-time river gauge metric measurement
type Reading struct {
	EventTime time.Time `json:"time"`
	Value     float32   `json:"value"`
}

// Snapshot is a set of measurements for a particular gauge station
type Snapshot struct {
	Station        Station   `json:"station"`
	Readings       []Reading `json:"readings"`
	TraceID        string    `json:"trace_id"`
	ProcessingTime time.Time `json:"processed_time"`
}
