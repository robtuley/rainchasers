package main

import (
	"testing"
)

func TestUpdatingFromAStation(t *testing.T) {
	readings, err := getReadings("http://apps.sepa.org.uk/database/riverlevels/116011-SG.csv")

	if err != nil {
		t.Error("Update stations error", err)
	}
	if len(readings) < 50 {
		t.Error("Not enough readings found", len(readings))
	}

	for i, u := range readings {
		if u.DateTime.IsZero() {
			t.Error("No DateTime", i)
		}
	}
}
