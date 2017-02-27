package osgridconverter

import (
	"math"
)

// Vector3d is a three dimensional vector whose properties
// contains spacial information about its location on the
// x-y-z axes of a three dimensional cartesian coordiantes system.
type Vector3d struct {
	x, y, z float64
}

// ToLatLon converts the Vector3d cartesian coordinates to
// latitude longitude according to the WGS84 datum.
// It returns a lat/lon coordinates struct.
func (v *Vector3d) ToLatLon(datum Datum) Coordinates {
	aa := datum.a
	bb := datum.b

	e2 := (aa*aa - bb*bb) / (aa * aa) // 1st eccentricity squared
	ε2 := (aa*aa - bb*bb) / (bb * bb) // 2nd eccentricity squared
	p := math.Sqrt(v.x*v.x + v.y*v.y) // distance from minor axis
	R := math.Sqrt(p*p + v.z*v.z)     // polar radius

	// parametric latitude
	tanβ := (bb * v.z) / (aa * p) * (1 + ε2*b/R)
	sinβ := tanβ / math.Sqrt(1+tanβ*tanβ)
	cosβ := sinβ / tanβ

	// geodetic latitude
	φ := math.Atan2(v.z+ε2*bb*sinβ*sinβ*sinβ, p-e2*aa*cosβ*cosβ*cosβ)

	// longitude
	λ := math.Atan2(v.y, v.x)

	coords := Coordinates{
		Lat: toDegrees(φ),
		Lon: toDegrees(λ),
	}

	return coords
}
