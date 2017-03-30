package main

import (
	"math"
	"testing"
)

func TestDiscoveringStations(t *testing.T) {
	const ε = 0.0001

	stations, err := discover()
	if err != nil {
		t.Error("Discover stations error", err)
	}

	if len(stations) < 200 {
		t.Error("Not enough stations found", len(stations))
	}

	nMissingNames := 0
	nMissingRiverNames := 0
	nMissingLat := 0
	nMissingLg := 0
	for i, s := range stations {
		if len(s.DataURL) < 5 {
			t.Error("No data URL", i, s)
		}
		if len(s.HumanURL) < 5 {
			t.Error("No human URL", i, s)
		}
		if len(s.Type) < 3 {
			t.Error("No type", i, s)
		}
		if len(s.Unit) < 1 {
			t.Error("No units", i, s)
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

	if nMissingNames > len(stations)/4 {
		t.Error("Too many missing names", nMissingNames, len(stations))
	}
	if nMissingRiverNames > 3*len(stations)/4 {
		t.Error("Too many missing river names", nMissingRiverNames, len(stations))
	}
	if nMissingLat > len(stations)/4 {
		t.Error("Too many missing lat", nMissingLat, len(stations))
	}
	if nMissingLg > len(stations)/4 {
		t.Error("Too many missing lg", nMissingLg, len(stations))
	}
}
