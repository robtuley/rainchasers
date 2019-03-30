package river

import (
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
)

// LatLng is a geographic location
type LatLng struct {
	Lat float32 `firestore:"lat" yaml:"lat"`
	Lng float32 `firestore:"lng" yaml:"lng"`
}

// Grade is the section difficultly
type Grade struct {
	Human   string  `firestore:"text" yaml:"text"`
	Average float32 `firestore:"value" yaml:"value"`
	Max     float32 `firestore:"max,omitempty" yaml:"max,omitempty"`
}

// Calibration is a referenced gauge related to a section
type Calibration struct {
	URL         string  `firestore:"-" yaml:"data_url"`
	Description string  `firestore:"desc" yaml:"desc"`
	Scrape      float32 `firestore:"scrape,omitempty" yaml:"scrape,omitempty"`
	Low         float32 `firestore:"low,omitempty" yaml:"low,omitempty"`
	Medium      float32 `firestore:"medium,omitempty" yaml:"medium,omitempty"`
	High        float32 `firestore:"high,omitempty" yaml:"high,omitempty"`
	Huge        float32 `firestore:"huge,omitempty" yaml:"huge,omitempty"`
	TooHigh     float32 `firestore:"toohigh,omitempty" yaml:"toohigh,omitempty"`
}

// Section is a paddleable river section
type Section struct {
	UUID        string        `firestore:"uuid" yaml:"uuid"`
	Slug        string        `firestore:"slug"` // derived from filename not yaml
	SectionName string        `firestore:"section" yaml:"section"`
	RiverName   string        `firestore:"river" yaml:"river"`
	KM          float32       `firestore:"km" yaml:"km"`
	Grade       Grade         `firestore:"grade" yaml:"grade"`
	Putin       LatLng        `firestore:"putin" yaml:"putin"`
	Takeout     LatLng        `firestore:"takeout" yaml:"takeout"`
	Description string        `firestore:"desc" yaml:"desc"`
	Directions  string        `firestore:"directions" yaml:"directions"`
	Measures    []Calibration `firestore:"-" yaml:"measures"`
}

// Checksum returns a comparable hash of the section
func (s Section) Checksum() string {
	b := sha1.New()
	gob.NewEncoder(b).Encode(s)
	return hex.EncodeToString(b.Sum(nil))
}
