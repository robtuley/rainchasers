package main

import (
	"testing"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
)

func TestUpdatingFromAStation(t *testing.T) {
	d := daemon.New("test", time.Second*60)
	readings, err := getReadings(d, "http://apps.sepa.org.uk/database/riverlevels/116011-SG.csv")

	if err != nil {
		t.Error("Update stations error", err)
	}
	if len(readings) < 50 {
		t.Error("Not enough readings found", len(readings))
	}

	for i, u := range readings {
		if u.EventTime.IsZero() {
			t.Error("No EventTime", i)
		}
	}
}
