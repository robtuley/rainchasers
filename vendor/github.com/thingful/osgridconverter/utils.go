package osgridconverter

import (
	"math"
)

// toDegrees converts radians to numeric degrees.
func toDegrees(input float64) float64 {
	return input * 180 / math.Pi
}

// toRadians converts numeric degrees to radians.
func toRadians(input float64) float64 {
	return input * math.Pi / 180
}
