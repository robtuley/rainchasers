package gauge

import (
	"regexp"
	"testing"
	"time"
)

func TestIdGeneration(t *testing.T) {
	validHex := regexp.MustCompile(`^[a-f0-9]+$`)

	s := Station{
		DataURL:   "http://environment.data.gov.uk/flood-monitoring/id/measures/1029TH-level-downstage-i-15_min-mASD",
		HumanURL:  "http://environment.data.gov.uk/flood-monitoring/id/stations/1029TH",
		Name:      "Bourton Dickler",
		RiverName: "Dikler",
		Lat:       51.874767,
		Lg:        -1.740083,
		Type:      "level",
		Unit:      "metre",
	}

	timestamp, _ := time.Parse(time.RFC3339, "2016-01-01T10:30:00Z")
	r := Reading{
		DateTime: timestamp,
		Value:    -0.14,
	}

	metricID := MetricID(s.DataURL)
	if !validHex.MatchString(metricID) {
		t.Error("MetricID ot valid hex string", metricID)
	}

	insertID := InsertID(s, r)
	if !validHex.MatchString(insertID) {
		t.Error("InsertID not valid hex string", insertID)
	}
	r.DateTime, _ = time.Parse(time.RFC3339, "2016-01-01T11:30:00Z")
	if insertID == InsertID(s, r) {
		t.Error("insertID does not change on timestamp", insertID, InsertID(s, r))
	}
}
