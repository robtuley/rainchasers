package main

import (
	"errors"
	"strconv"
	"strings"

	"github.com/thingful/osgridconverter"
)

func gridRef2LatLg(gr string) (float32, float32, error) {
	e, n, err := parseGridRef(gr)
	if err != nil {
		return 0.0, 0.0, err
	}
	coords, err := osgridconverter.ConvertToLatLon(float64(e), float64(n), osgridconverter.WGS84)
	if err != nil {
		return 0.0, 0.0, err
	}

	return float32(coords.Lat), float32(coords.Lon), nil
}

// Converts standard grid reference ('NS4988662918') to fully numeric easting northing
func parseGridRef(gr string) (int, int, error) {
	gr = strings.ToUpper(gr)
	if len(gr) != 12 {
		return 0, 0, errors.New(gr + " is not 12 characters like NS4988662918")
	}

	e, err := strconv.Atoi(gr[2:7])
	if err != nil {
		return 0, 0, errors.New(gr + " eastings parse fail")
	}
	n, err := strconv.Atoi(gr[7:12])
	if err != nil {
		return 0, 0, errors.New(gr + " northing parse fail")
	}

	// get numeric values of letter references, mapping A->0, B->1, C->2, etc:
	r := []rune(gr)
	l1 := int(r[0] - []rune("A")[0])
	l2 := int(r[1] - []rune("A")[0])

	// shuffle down letters after 'I' since 'I' is not used in grid:
	if l1 > 7 {
		l1 = l1 - 1
	}
	if l2 > 7 {
		l2 = l2 - 1
	}

	// convert grid letters into 100km-square indexes from false origin (grid square SV)
	eMult := ((l1-2)%5)*5 + (l2 % 5)
	nMult := (19 - int(l1/5)*5) - int(l2/5.0)

	if eMult < 0 || eMult > 6 || nMult < 0 || nMult > 12 {
		return 0, 0, errors.New(gr[0:2] + " are invalid grid prefix letters")
	}
	return eMult*100000 + e, nMult*100000 + n, nil
}
