package main

import (
	"testing"
	"time"
)

func TestDownloadCSVParse(t *testing.T) {
	twoDaysAgo := time.Now().AddDate(0, 0, -2)
	snapshots, err := download(twoDaysAgo)
	if err != nil {
		t.Error("download error", err)
	}
	if len(snapshots) < 10000 {
		t.Error("Not enough readings found", len(snapshots))
	}

	for i, s := range snapshots {
		if len(s.MetricID) < 32 {
			t.Error("No Metric ID", i, s)
		}
		if s.DateTime.IsZero() {
			t.Error("No DateTime", i, s)
		}
	}
}
