package gauge

import (
	"crypto/sha1"
	"fmt"
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

// valid version 4 UUID (see RFC 4122, section 4.4)
func (s *Station) UUID() string {
	b := sha1.Sum([]byte(s.DataURL))
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

type Reading struct {
	DateTime time.Time
	Value    float32
}

type Snapshot struct {
	Station  Station
	Readings []Reading
}
