package main

import (
	"context"
	"testing"
)

func TestUpdatingFromAStation(t *testing.T) {
	readings, span := getReadings(context.Background(), "http://apps.sepa.org.uk/database/riverlevels/116011-SG.csv")

	if err := span.Err(); err != nil {
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
