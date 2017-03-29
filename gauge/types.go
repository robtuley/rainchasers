package gauge

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"strings"
	"time"
)

type Station struct {
	DataURL   string
	HumanURL  string
	Name      string
	RiverName string
	Lat       float32
	Lg        float32
	Type      string
	Unit      string
}

type Reading struct {
	DateTime time.Time
	Value    float32
}

type Snapshot struct {
	Station  Station
	Readings []Reading
}

func InsertID(s Station, r Reading) string {
	h := sha256.New()
	io.WriteString(h, r.DateTime.Format(time.RFC822)+strings.ToLower(s.DataURL))
	return hex.EncodeToString(h.Sum(nil))
}

func MetricID(dataURL string) string {
	h := sha256.New()
	io.WriteString(h, strings.ToLower(dataURL))
	return hex.EncodeToString(h.Sum(nil))
}
