package ea

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
)

func TestRecentReadings(t *testing.T) {
	d := daemon.New("example")
	readings, err := Recent(context.Background(), d)

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
	d := daemon.New("example")
	readings, err := Recent(context.Background(), d)
	if err != nil {
		t.Error("Update stations error", err)
	}

	stations, err := Discover(context.Background(), d)
	if err != nil {
		t.Error("Discover stations error", err)
	}

	nSkippedMetrics := 0
	for id, r := range readings {
		s, ok := stations[id]
		if !ok {
			nSkippedMetrics++
			continue
		}

		snap := gauge.Snapshot{
			Station:  s,
			Readings: []gauge.Reading{r},
		}
		err := snap.Encode(ioutil.Discard)
		if err != nil {
			t.Error("encoding snapshot error", err)
		}
	}
	if nSkippedMetrics > len(readings)/4 {
		t.Error("Too many skipped updates", nSkippedMetrics, len(readings))
	}
}
