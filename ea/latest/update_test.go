package main

import (
	"github.com/rainchasers/com.rainchasers.gauge/ea/discover"
	"github.com/rainchasers/com.rainchasers.gauge/gauge"
	"testing"
)

func TestUpdatingStations(t *testing.T) {
	readings, err := update()

	if err != nil {
		t.Error("Update stations error", err)
	}
	if len(readings) < 4000 {
		t.Error("Not enough readings found", len(readings))
	}

	for url, u := range readings {
		if len(url) < 32 {
			t.Error("Too short data URL", url)
		}
		if u.EventTime.IsZero() {
			t.Error("No EventTime", url)
		}
	}
}

func TestUpdatesAreForDiscoveredStations(t *testing.T) {
	readings, err := update()
	if err != nil {
		t.Error("Update stations error", err)
	}

	stations, err := discover.Stations()
	if err != nil {
		t.Error("Discover stations error", err)
	}

	nSkippedMetrics := 0
	for id, r := range readings {
		s, ok := stations[id]
		if !ok {
			nSkippedMetrics += 1
			continue
		}

		snap := gauge.Snapshot{
			Station:  s,
			Readings: []gauge.Reading{r},
		}
		_, err := snap.Encode()
		if err != nil {
			t.Error("encoding snapshot error", err)
		}
	}
	if nSkippedMetrics > len(readings)/4 {
		t.Error("Too many skipped updates", nSkippedMetrics, len(readings))
	}
}
