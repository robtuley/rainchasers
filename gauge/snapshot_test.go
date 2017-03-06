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
		DataURL:   "http://environment.data.gov.uk/flood-monitoring/id/measures/1029TH-level-downstage-i-15_min-mASD",
		HumanURL:  "http://environment.data.gov.uk/flood-monitoring/id/stations/1029TH",
		Name:      "Bourton Dickler",
		RiverName: "Dikler",
		Lat:       51.874767,
		Lg:        -1.740083,
		Type:      "level",
		Unit:      "metre",
		DateTime:  timestamp,
		Value:     -0.14,
	}

	insertID := snap.InsertID()
	metricID := snap.MetricID()

	var validHex = regexp.MustCompile(`^[a-f0-9]+$`)
	if !validHex.MatchString(insertID) {
		t.Error("InsertID not valid hex string", insertID)
	}
	if insertID != snap.InsertID() {
		t.Error("Non-deterministic insertID", insertID)
	}
	if !validHex.MatchString(metricID) {
		t.Error("MetricID ot valid hex string", insertID)
	}
	if metricID != snap.MetricID() {
		t.Error("Non-deterministic metricID", insertID)
	}

	snap.DataURL = "http://different"
	if insertID == snap.InsertID() {
		t.Error("insertID does not change on URL", insertID)
	}
	if metricID == snap.MetricID() {
		t.Error("metricID does not change on URL", insertID)
	}

	insertID = snap.InsertID()
	metricID = snap.MetricID()
	diffTimestamp, _ := time.Parse(time.RFC3339, "2016-01-01T11:30:00Z")
	snap.DateTime = diffTimestamp
	if insertID == snap.InsertID() {
		t.Error("insertID does not change on timestamp", insertID)
	}
	if metricID != snap.MetricID() {
		t.Error("metricID changes on timestamp", insertID)
	}
	if gauge.CalculateMetricID(snap.DataURL) != snap.MetricID() {
		t.Error("metricID different between calculate and embedded method", snap.MetricID(), gauge.CalculateMetricID(snap.DataURL))
	}

}
