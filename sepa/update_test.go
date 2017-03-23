package main

import (
	"testing"
)

func TestUpdatingFromAStation(t *testing.T) {
	updates, err := getReadings("http://apps.sepa.org.uk/database/riverlevels/116011-SG.csv")

	if err != nil {
		t.Error("Update stations error", err)
	}
	if len(updates) < 50 {
		t.Error("Not enough readings found", len(updates))
	}

	for i, u := range updates {
		if u.DateTime.IsZero() {
			t.Error("No DateTime", i)
		}
	}
}
