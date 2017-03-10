package main

import (
	"testing"
)

func TestUpdatingStations(t *testing.T) {
	updates, err := update()

	if err != nil {
		t.Error("Update stations error", err)
	}
	if len(updates) < 4000 {
		t.Error("Not enough readings found", len(updates))
	}

	for i, u := range updates {
		if len(u.MetricID) < 32 {
			t.Error("No Metric ID", i, u)
		}
		if u.DateTime.IsZero() {
			t.Error("No DateTime", i, u)
		}
	}
}

func TestUpdatesAreForDiscoveredStations(t *testing.T) {
	updates, err := update()
	if err != nil {
		t.Error("Update stations error", err)
	}

	snapshots, err := discover()
	if err != nil {
		t.Error("Discover stations error", err)
	}

	nSkippedMetrics := 0
	for _, u := range updates {
		_, ok := snapshots[u.MetricID]
		if !ok {
			nSkippedMetrics += 1
		}
	}
	if nSkippedMetrics > len(updates)/4 {
		t.Error("Too many skipped updates", nSkippedMetrics, len(updates))
	}
}
