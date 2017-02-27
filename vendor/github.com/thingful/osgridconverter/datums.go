package osgridconverter

// Datum is a model of the earth that is used in mapping.
// For more informations see https://en.wikipedia.org/wiki/Geodetic_datum
type Datum struct {
	a  float64 // major axis
	b  float64 // minor axis
	f  float64 // flattening
	tx float64 // m
	ty float64 // m
	tz float64 // m
	rx float64 // sec
	ry float64 // sec
	rz float64 // sec
	s  float64 // ppm
}

var (
	// WGS84 datum
	WGS84 = Datum{
		a:  6378137,
		b:  6356752.31425,
		f:  0.0,
		tx: 0.0,
		ty: 0.0,
		tz: 0.0,
		rx: 0.0,
		ry: 0.0,
		rz: 0.0,
		s:  0.0,
	}

	// NAD83 datum. Functionally ≡ WGS84
	// uses GRS80 ellipsoid parameters
	NAD83 = Datum{
		a:  6378137,
		b:  6356752.31414,
		f:  1 / 298.257222101,
		tx: 1.004,
		ty: -1.910,
		tz: -0.515,
		rx: 0.0267,
		ry: 0.00034,
		rz: 0.011,
		s:  -0.0015,
	}

	// OSGB36 datum
	// uses Airy1830 ellipsoid parameters
	OSGB36 = Datum{
		a:  6377563.396,
		b:  6356256.909,
		f:  1 / 299.3249646,
		tx: -446.448,
		ty: 125.157,
		tz: -542.060,
		rx: -0.1502,
		ry: -0.2470,
		rz: -0.8421,
		s:  20.4894,
	}

	// ED50 datum
	// uses Intl1924 ellipsoid parameters
	ED50 = Datum{
		a:  6378388,
		b:  6356911.946,
		f:  1 / 297,
		tx: 89.5,
		ty: 93.8,
		tz: 123.1,
		rx: 0.0,
		ry: 0.0,
		rz: 0.156,
		s:  -1.2,
	}

	// Irl1975 datum
	// uses AiryModified ellipsoid parameters
	Irl1975 = Datum{
		a:  6377340.189,
		b:  6356034.448,
		f:  1 / 299.3249646,
		tx: -482.530,
		ty: 130.596,
		tz: -564.557,
		rx: -1.042,
		ry: -0.214,
		rz: -0.631,
		s:  -8.150,
	}

	// TokyoJapan datum
	// uses Bessel1841 ellipsoid parameters
	TokyoJapan = Datum{
		a:  6377397.155,
		b:  6356078.963,
		f:  1 / 299.152815351,
		tx: 148,
		ty: -507,
		tz: -685,
		rx: 0,
		ry: 0,
		rz: 0,
		s:  0,
	}
)
