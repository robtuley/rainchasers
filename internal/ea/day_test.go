package ea

import (
	"context"
	"testing"
	"time"
)

func TestDayCSVParse(t *testing.T) {
	twoDaysAgo := time.Now().AddDate(0, 0, -2)
	snapshots, span := Day(context.Background(), twoDaysAgo)
	if err := span.Err(); err != nil {
		t.Error("download error", err)
	}
	if len(snapshots) < 2000 {
		t.Error("Not enough gauges found", len(snapshots))
	}

	for id, snaps := range snapshots {
		if len(id) < 32 {
			t.Error("Bad Data URL", id)
		}
		for i, s := range snaps {
			if s.EventTime.IsZero() {
				t.Error("No EventTime", id, i, s)
			}
		}
	}
}
