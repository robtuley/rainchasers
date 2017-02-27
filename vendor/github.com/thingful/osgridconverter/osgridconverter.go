// Package osgridconverter contains utility functions to convert
// Ordnance Survey grid references to latitude/longitude coordinates.
// By Default the library will return latitude and longitude according
// to the OSGB36 datum which is generally used for GPS systems
package osgridconverter

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

const (
	a  = 6377563.396        // Airy 1830 major & minor semi-axes
	b  = 6356256.909        // Airy 1830 major & minor semi-axes
	f0 = 0.9996012717       // NatGrid scale factor on central meridian
	φ0 = 49 * math.Pi / 180 // NatGrid true origin
	λ0 = -2 * math.Pi / 180 // NatGrid true origin
	n0 = -100000            // northing of true origin, metres
	e0 = 400000             // easting of true origin, metres
	e2 = 1 - (b*b)/(a*a)    // eccentricity squared
	n  = (a - b) / (a + b)  // n
	n2 = n * n              // n²
	n3 = n * n * n          // n³
)

var (
	// ErrInvalidOSGridPoints is an error that can be returned when OS
	// grid points passed to the converter function have negative values
	ErrInvalidOSGridPoints = errors.New("osgridconverter: Invalid arguments. Easting and Northing coordinates should be positive float64")

	// ErrInvalidLat is an error that can be returned when the latitude
	// value passed to the converter function is lower than -90 or highter than +90
	ErrInvalidLat = errors.New("osgridconverter: Latitude values must be between -90 and +90")

	// ErrInvalidLon is an error that can be returned when the longitude
	// value passed to the converter function is lower than -180 or highter than +180
	ErrInvalidLon = errors.New("osgridconverter: Longitude values must be between -180 and +180")
)

// Coordinates struct holds the latitude and longitude coordinates
type Coordinates struct {
	Lat float64
	Lon float64
}

// OsGrid struct holds the easting and northing grid points
type OsGrid struct {
	Easting  float64
	Northing float64
}

// ConvertToLatLon converts Ordnance Survey grid reference easting and northing
// coordinates to latitude and longitude according to the WGS-84 ellipsoidal model.
// Easting and Northing arguments should be numeric references in metres (eg 438700, 114800).
// It returns a struct containing latitude and longitude coordinates as float64 type
// or an error if the arguments passed in are out of bounds
func ConvertToLatLon(easting, northing float64, datum Datum) (*Coordinates, error) {
	c := Coordinates{}

	// validate input
	if easting < 0 || northing < 0 {
		return &c, ErrInvalidOSGridPoints
	}

	φ := φ0
	M := float64(0)

	for northing-n0-M >= 0.00001 {
		φ = (northing-n0-M)/(a*f0) + φ
		Ma := (1 + n + (5/4)*n2 + (5/4)*n3) * (φ - φ0)
		Mb := (3*n + 3*n*n + (21/8)*n3) * math.Sin(φ-φ0) * math.Cos(φ+φ0)
		Mc := ((15/8)*n2 + (15/8)*n3) * math.Sin(2*(φ-φ0)) * math.Cos(2*(φ+φ0))
		Md := (35 / 24) * n3 * math.Sin(3*(φ-φ0)) * math.Cos(3*(φ+φ0))
		M = b * f0 * (Ma - Mb + Mc - Md) // meridional arc
	}

	cosφ := math.Cos(φ)
	sinφ := math.Sin(φ)
	ν := a * f0 / math.Sqrt(1-e2*sinφ*sinφ)                // nu = transverse radius of curvature
	ρ := a * f0 * (1 - e2) / math.Pow(1-e2*sinφ*sinφ, 1.5) // rho = meridional radius of curvature
	η2 := ν/ρ - 1

	tanφ := math.Tan(φ)
	tan2φ := tanφ * tanφ
	tan4φ := tan2φ * tan2φ
	tan6φ := tan4φ * tan2φ
	secφ := 1 / cosφ
	ν3 := ν * ν * ν
	ν5 := ν3 * ν * ν
	ν7 := ν5 * ν * ν
	VII := tanφ / (2 * ρ * ν)
	VIII := tanφ / (24 * ρ * ν3) * (5 + 3*tan2φ + η2 - 9*tan2φ*η2)
	IX := tanφ / (720 * ρ * ν5) * (61 + 90*tan2φ + 45*tan4φ)
	X := secφ / ν
	XI := secφ / (6 * ν3) * (ν/ρ + 2*tan2φ)
	XII := secφ / (120 * ν5) * (5 + 28*tan2φ + 24*tan4φ)
	XIIA := secφ / (5040 * ν7) * (61 + 662*tan2φ + 1320*tan4φ + 720*tan6φ)

	dE := (easting - e0)
	dE2 := dE * dE
	dE3 := dE2 * dE
	dE4 := dE2 * dE2
	dE5 := dE3 * dE2
	dE6 := dE4 * dE2
	dE7 := dE5 * dE2
	φ = φ - VII*dE2 + VIII*dE4 - IX*dE6
	λ := λ0 + X*dE - XI*dE3 + XII*dE5 - XIIA*dE7

	c.Lat = toDegrees(φ)
	c.Lon = toDegrees(λ)

	if datum == OSGB36 {
		return &c, nil
	}

	// convert to Vector3d
	vect := toCartesian(c, OSGB36)

	// convert back to coordinates
	cc := vect.ToLatLon(datum)

	return &cc, nil
}

// ConvertToNorthingEasting converts latitude and longitude to
// Ordnance Survey grid reference northing and easting.
// It returns a struct containing easting and northing coordinates as float64 type
// or an error if the arguments passed in are out of bounds
func ConvertToNorthingEasting(lat, lon float64) (*OsGrid, error) {
	o := OsGrid{}

	// validate input
	if lat < -90 || lat > 90 {
		return &o, ErrInvalidLat
	}

	if lon < -180 || lon > 180 {
		return &o, ErrInvalidLon
	}

	φ := toRadians(lat)
	λ := toRadians(lon)

	cosφ := math.Cos(φ)
	sinφ := math.Sin(φ)
	ν := a * f0 / math.Sqrt(1-e2*sinφ*sinφ)
	ρ := a * f0 * (1 - e2) / math.Pow(1-e2*sinφ*sinφ, 1.5)
	η2 := ν/ρ - 1

	Ma := (1 + n + (5/4)*n2 + (5/4)*n3) * (φ - φ0)
	Mb := (3*n + 3*n*n + (21/8)*n3) * math.Sin(φ-φ0) * math.Cos(φ+φ0)
	Mc := ((15/8)*n2 + (15/8)*n3) * math.Sin(2*(φ-φ0)) * math.Cos(2*(φ+φ0))
	Md := (35 / 24) * n3 * math.Sin(3*(φ-φ0)) * math.Cos(3*(φ+φ0))
	M := b * f0 * (Ma - Mb + Mc - Md)

	cos3φ := cosφ * cosφ * cosφ
	cos5φ := cos3φ * cosφ * cosφ
	tan2φ := math.Tan(φ) * math.Tan(φ)
	tan4φ := tan2φ * tan2φ

	I := M + n0
	II := (ν / 2) * sinφ * cosφ
	III := (ν / 24) * sinφ * cos3φ * (5 - tan2φ + 9*η2)
	IIIA := (ν / 720) * sinφ * cos5φ * (61 - 58*tan2φ + tan4φ)
	IV := ν * cosφ
	V := (ν / 6) * cos3φ * (ν/ρ - tan2φ)
	VI := (ν / 120) * cos5φ * (5 - 18*tan2φ + tan4φ + 14*η2 - 58*tan2φ*η2)

	Δλ := λ - λ0
	Δλ2 := Δλ * Δλ
	Δλ3 := Δλ2 * Δλ
	Δλ4 := Δλ3 * Δλ
	Δλ5 := Δλ4 * Δλ
	Δλ6 := Δλ5 * Δλ

	northingVal := I + II*Δλ2 + III*Δλ4 + IIIA*Δλ6
	northingVal, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", northingVal), 64) // truncate after 3 decimal positions
	o.Northing = northingVal

	eastingVal := e0 + IV*Δλ + V*Δλ3 + VI*Δλ5
	eastingVal, _ = strconv.ParseFloat(fmt.Sprintf("%.3f", eastingVal), 64) // truncate after 3 decimal positions
	o.Easting = eastingVal

	return &o, nil
}

// toCartesian converts lat/lon coordinates to cartesian x/y/z coordinates.
// It returns a vector representing a lat/lon point on x-y-z axes in metres
// from the earth centre.
func toCartesian(coords Coordinates, datum Datum) Vector3d {

	aa := datum.a
	bb := datum.b

	φ := toRadians(coords.Lat)
	λ := toRadians(coords.Lon)
	sinφ := math.Sin(φ)
	cosφ := math.Cos(φ)
	sinλ := math.Sin(λ)
	cosλ := math.Cos(λ)

	eSq := (aa*aa - bb*bb) / (aa * aa) // make a and b const datum properties
	ν := aa / math.Sqrt(1-eSq*sinφ*sinφ)
	x := ν * cosφ * cosλ
	y := ν * cosφ * sinλ
	z := (1 - eSq) * ν * sinφ

	// apply Helmert transform
	// move to another function
	rx := toRadians((datum.rx * -1) / 3600) // normalise seconds to radians
	ry := toRadians((datum.ry * -1) / 3600) // normalise seconds to radians
	rz := toRadians((datum.rz * -1) / 3600) // normalise seconds to radians
	s1 := (datum.s*-1)/1e6 + 1              // normalise ppm to (s+1)

	// apply transform
	x2 := (datum.tx * -1) + x*s1 - y*rz + z*ry
	y2 := (datum.ty * -1) + x*rz + y*s1 - z*rx
	z2 := (datum.tz * -1) - x*ry + y*rx + z*s1

	// instantiate new vector
	vect := Vector3d{x: x2, y: y2, z: z2}

	return vect
}
