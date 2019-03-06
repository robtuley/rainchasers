package gauge

import (
	"crypto/sha1"
	"fmt"
	"time"
)

// Station is a specific location that can measure one or more metrics
type Station struct {
	DataURL   string  `json:"data_url"`
	HumanURL  string  `json:"human_url"`
	Name      string  `json:"name"`
	RiverName string  `json:"river"`
	Lat       float32 `json:"lat"`
	Lg        float32 `json:"lng"`
	Type      string  `json:"type"`
	Unit      string  `json:"unit"`
}

// UUID provides a deterministic valid version 4 UUID based on the data URL
func (s Station) UUID() string {
	b := sha1.Sum([]byte(s.DataURL))
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
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
	ProcessingTime time.Time `json:"processed_time"`
}
