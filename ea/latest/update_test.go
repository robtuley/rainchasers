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
