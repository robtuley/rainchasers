package main

import (
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
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
		ref, ok := snapshots[u.MetricID]
		if !ok {
			nSkippedMetrics += 1
			continue
		}

		s := ref.Apply(u)
		_, err := gauge.EncodeSnapshot(s)
		if err != nil {
			t.Error("encoding snapshot error", err)
		}
	}
	if nSkippedMetrics > len(updates)/4 {
		t.Error("Too many skipped updates", nSkippedMetrics, len(updates))
	}
}
