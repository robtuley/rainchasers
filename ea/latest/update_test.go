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

	for url, u := range updates {
		if len(url) < 32 {
			t.Error("Too short data URL", url)
		}
		if u.DateTime.IsZero() {
			t.Error("No DateTime", url)
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
	for id, u := range updates {
		ref, ok := snapshots[id]
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
