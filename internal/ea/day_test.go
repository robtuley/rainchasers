package ea

import (
	"testing"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
)

func TestDayCSVParse(t *testing.T) {
	d := daemon.New("example", 2*time.Minute)

	twoDaysAgo := time.Now().AddDate(0, 0, -2)
	snapshots, err := Day(d, twoDaysAgo)
	if err != nil {
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
