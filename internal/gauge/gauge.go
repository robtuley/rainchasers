package gauge

import (
	"time"
)

// Station is a specific location that can measure one or more metrics
type Station struct {
	DataURL   string  `firestore:"data_url" json:"data_url"`
	AliasURL  string  `firestore:"alias_url" json:"alias_url"`
	HumanURL  string  `firestore:"human_url" json:"human_url"`
	Name      string  `firestore:"name" json:"name"`
	RiverName string  `firestore:"river" json:"river"`
	Lat       float32 `firestore:"lat" json:"lat"`
	Lg        float32 `firestore:"lng" json:"lng"`
	Type      string  `firestore:"type" json:"type"`
	Unit      string  `firestore:"unit" json:"unit"`
}

// Reading is a point-in-time river gauge metric measurement
type Reading struct {
	EventTime time.Time `firestore:"time" json:"time"`
	Value     float32   `firestore:"value" json:"value"`
}

// Snapshot is a set of measurements for a particular gauge station
type Snapshot struct {
	Station       Station   `json:"station"`
	Readings      []Reading `json:"readings"`
	CorrelationID string    `json:"correlation_id"`
	CausationID   string    `json:"causation_id"`
	ProcessedTime time.Time `json:"processed_time"`
}
