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
	if len(snapshots) < 2000 {
		t.Error("Not enough gauges found", len(snapshots))
	}

	for id, snaps := range snapshots {
		if len(id) < 32 {
			t.Error("Bad Data URL", id)
		}
		for i, s := range snaps {
			if s.DateTime.IsZero() {
				t.Error("No DateTime", id, i, s)
			}
		}
	}
}
