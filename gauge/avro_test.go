package gauge_test

import (
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"testing"
	"time"
)

func TestSnapshotEncodeDecode(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-01-01T10:30:00Z")

	before := gauge.Snapshot{
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

	bb, err := gauge.EncodeSnapshot(before)
	if err != nil {
		t.Error(err)
	}

	after, err := gauge.DecodeSnapshot(bb)
	if err != nil {
		t.Error(err)
	}

	// check fields individually (not using reflect.DeepEqual as
	// some custom compare needed for the dates)
	if before.DataURL != after.DataURL {
		t.Error("Url mis-match", after)
	}
	if before.HumanURL != after.HumanURL {
		t.Error("Station Url mis-match", after)
	}
	if before.Name != after.Name {
		t.Error("Name mis-match", after)
	}
	if before.RiverName != after.RiverName {
		t.Error("River name mis-match", after)
	}
	if before.Lat != after.Lat {
		t.Error("Url mis-match", after)
	}
	if before.Lg != after.Lg {
		t.Error("Lg mis-match", after)
	}
	if before.Type != after.Type {
		t.Error("Type mis-match", after)
	}
	if before.Unit != after.Unit {
		t.Error("Unit mis-match", after)
	}
	if before.DateTime.Unix() != after.DateTime.Unix() {
		t.Error("Timestamp mis-match", before.DateTime.Unix(), after.DateTime.Unix(), after)
	}
	if before.Value != after.Value {
		t.Error("Value mis-match", after)
	}
}

func TestSnapshotUpdatesEncodeDecode(t *testing.T) {
	timestamp, _ := time.Parse(time.RFC3339, "2016-01-01T10:30:00Z")
	var before []gauge.SnapshotUpdate

	before = append(before, gauge.SnapshotUpdate{
		DateTime: timestamp.Add(time.Second),
		Value:    1.23,
	})
	before = append(before, gauge.SnapshotUpdate{
		DateTime: timestamp.Add(time.Second * 10),
		Value:    4.56,
	})

	bb, err := gauge.EncodeSnapshotUpdates("123456", before)
	if err != nil {
		t.Error(err)
	}

	url, after, err := gauge.DecodeSnapshotUpdates(bb)
	if err != nil {
		t.Error(err)
	}

	if len(after) != len(before) {
		t.Error("length mismatch", len(before), len(after))
		return
	}

	if url != "123456" {
		t.Error("Data URL mismatch", url)
	}

	for i, b := range before {
		a := after[i]
		if b.DateTime.Unix() != a.DateTime.Unix() {
			t.Error("Timestamp mis-match", i, b.DateTime.Unix(), a.DateTime.Unix())
		}
		if b.Value != a.Value {
			t.Error("Value mis-match", i, b.Value, a.Value)
		}
	}
}
