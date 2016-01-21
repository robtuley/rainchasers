package gauge_test

import (
	"regexp"
	"testing"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/gauge"
)

func TestSnapshotInsertIdGeneration(t *testing.T) {
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

	var validHex = regexp.MustCompile(`^[a-f0-9]+$`)
	if !validHex.MatchString(insertId) {
		t.Error("Not valid hex string", insertId)
	}
	if insertId != snap.InsertId() {
		t.Error("Non-deterministic insertId", insertId)
	}

	snap.Url = "http://different"
	if insertId == snap.InsertId() {
		t.Error("insertId does not change on URL", insertId)
	}

	insertId = snap.InsertId()
	diffTimestamp, _ := time.Parse(time.RFC3339, "2016-01-01T11:30:00Z")
	snap.DateTime = diffTimestamp
	if insertId == snap.InsertId() {
		t.Error("insertId does not change on timestamp", insertId)
	}
}
