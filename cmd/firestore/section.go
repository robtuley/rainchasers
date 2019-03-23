package main

import (
	"crypto/sha1"
	"encoding/gob"
	"encoding/hex"
)

// LatLng is a geographic location
type LatLng struct {
	Lat float32 `firestore:"lat" json:"lat"`
	Lng float32 `firestore:"lng" json:"lng"`
}

// Grade is the section difficultly
type Grade struct {
	Human   string   `firestore:"text" json:"text"`
	Average float32  `firestore:"value" json:"value"`
	Max     *float32 `firestore:"max,omitempty" json:"max,omitempty"`
}

// Calibration is a referenced gauge related to a section
type Calibration struct {
	DataURL     string   `firestore:"data_url" json:"data_url"`
	Description string   `firestore:"desc" json:"desc"`
	Scrape      *float32 `firestore:"scrape,omitempty" json:"scrape,omitempty"`
	Low         *float32 `firestore:"low,omitempty" json:"low,omitempty"`
	Medium      *float32 `firestore:"medium,omitempty" json:"medium,omitempty"`
	High        *float32 `firestore:"high,omitempty" json:"high,omitempty"`
	Huge        *float32 `firestore:"huge,omitempty" json:"huge,omitempty"`
	TooHigh     *float32 `firestore:"toohigh,omitempty" json:"toohigh,omitempty"`
}

// Section is a paddleable river section
type Section struct {
	UUID        string        `firestore:"uuid" json:"uuid"`
	SectionName string        `firestore:"section" json:"section"`
	RiverName   string        `firestore:"river" json:"river"`
	KM          float32       `firestore:"km" json:"km"`
	Grade       Grade         `firestore:"grade" json:"grade"`
	Putin       LatLng        `firestore:"putin" json:"putin"`
	Takeout     LatLng        `firestore:"takeout" json:"takeout"`
	Description string        `firestore:"desc" json:"desc"`
	Directions  string        `firestore:"directions" json:"directions"`
	Measures    []Calibration `firestore:"measures" json:"measures"`
}

// Checksum returns a comparable hash of the section
func (s Section) Checksum() string {
	b := sha1.New()
	gob.NewEncoder(b).Encode(s)
	return hex.EncodeToString(b.Sum(nil))
}
