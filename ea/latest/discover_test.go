package main

import (
	"math"
	"testing"
)

const ε = 0.0001

func TestDiscoveringStations(t *testing.T) {
	snapshots, err := discover()
	if err != nil {
		t.Error("Discover stations error", err)
	}

	if len(snapshots) < 4000 {
		t.Error("Not enough stations found", len(snapshots))
	}

	nMissingNames := 0
	nMissingRiverNames := 0
	nMissingLat := 0
	nMissingLg := 0
	for id, s := range snapshots {
		if id != s.DataURL {
			t.Error("Data URL not mapped", id, s.DataURL)
		}
		if len(s.DataURL) < 5 {
			t.Error("No data URL", id, s)
		}
		if len(s.HumanURL) < 5 {
			t.Error("No human URL", id, s)
		}
		if len(s.Type) < 3 {
			t.Error("No type", id, s)
		}
		if len(s.Unit) < 1 {
			t.Error("No units", id, s)
		}
		if len(s.Name) < 3 {
			nMissingNames += 1
		}
		if len(s.RiverName) < 3 {
			nMissingRiverNames += 1
		}
		if δ := math.Abs(float64(s.Lat)); δ < ε {
			nMissingLat += 1
		}
		if δ := math.Abs(float64(s.Lg)); δ < ε {
			nMissingLg += 1
		}

	}

	if nMissingNames > len(snapshots)/4 {
		t.Error("Too many missing names", nMissingNames, len(snapshots))
	}
	if nMissingRiverNames > 3*len(snapshots)/4 {
		t.Error("Too many missing river names", nMissingRiverNames, len(snapshots))
	}
	if nMissingLat > len(snapshots)/4 {
		t.Error("Too many missing lat", nMissingLat, len(snapshots))
	}
	if nMissingLg > len(snapshots)/4 {
		t.Error("Too many missing lg", nMissingLg, len(snapshots))
	}
}
