package gauge_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
)

func TestSnapshotIdGeneration(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-01-01T10:30:00Z")

	snap := gauge.Snapshot{
		Url:        "http://environment.data.gov.uk/flood-monitoring/id/measures/1029TH-level-downstage-i-15_min-mASD",
		StationUrl: "http://environment.data.gov.uk/flood-monitoring/id/stations/1029TH",
		Name:       "Bourton Dickler",
		RiverName:  "Dikler",
		Lat:        51.874767,
		Lg:         -1.740083,
		Type:       "level",
		Unit:       "metre",
		DateTime:   timestamp,
		Value:      -0.14,
	}

	insertId := snap.InsertId()
	metricId := snap.MetricId()

	var validHex = regexp.MustCompile(`^[a-f0-9]+$`)
	if !validHex.MatchString(insertId) {
		t.Error("InsertId not valid hex string", insertId)
	}
	if insertId != snap.InsertId() {
		t.Error("Non-deterministic insertId", insertId)
	}
	if !validHex.MatchString(metricId) {
		t.Error("MetricId ot valid hex string", insertId)
	}
	if metricId != snap.MetricId() {
		t.Error("Non-deterministic metricId", insertId)
	}

	snap.Url = "http://different"
	if insertId == snap.InsertId() {
		t.Error("insertId does not change on URL", insertId)
	}
	if metricId == snap.MetricId() {
		t.Error("metricId does not change on URL", insertId)
	}

	insertId = snap.InsertId()
	metricId = snap.MetricId()
	diffTimestamp, _ := time.Parse(time.RFC3339, "2016-01-01T11:30:00Z")
	snap.DateTime = diffTimestamp
	if insertId == snap.InsertId() {
		t.Error("insertId does not change on timestamp", insertId)
	}
	if metricId != snap.MetricId() {
		t.Error("metricId changes on timestamp", insertId)
	}
}
