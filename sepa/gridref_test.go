package main

import (
	"math"
	"testing"
)

func TestParsingGridReference(t *testing.T) {
	e, n, err := parseGridRef("NO1160525332")
	if err != nil {
		t.Error(err, e, n)
		return
	}
	if e != 311605 {
		t.Error("Incorrect easting from grid ref", e, n)
	}
	if n != 725332 {
		t.Error("Incorrect northing from grid ref", e, n)
	}

	_, _, err = parseGridRef("ZX1160525332")
	if err == nil {
		t.Error("Invalid letters prefix accepted by parser")
	}
}

func TestGridReferenceConverstion(t *testing.T) {
	const ε = 0.0001

	lat, lg, err := gridRef2LatLg("NO1160525332")
	if err != nil {
		t.Error(err, lat, lg)
		return
	}

	if δ := math.Abs(float64(lat) - 56.411915); δ > ε {
		t.Error("Incorrect lat from grid ref", lat)
	}
	if δ := math.Abs(float64(lg) + 3.434206); δ > ε {
		t.Error("Incorrect lg from grid ref", lg)
	}

	_, _, err = gridRef2LatLg("ZX1160525332")
	if err == nil {
		t.Error("Invalid letters prefix accepted by parser")
	}
}
