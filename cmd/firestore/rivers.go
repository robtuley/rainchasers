package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
	"github.com/rainchasers/report"
)

//  Catalogue from https://app.rainchasers.com/catalogue.json

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
	DataURL     string          `firestore:"data_url" json:"data_url"`
	Description string          `firestore:"desc" json:"desc"`
	Scrape      *float32        `firestore:"scrape,omitempty" json:"scrape,omitempty"`
	Low         *float32        `firestore:"low,omitempty" json:"low,omitempty"`
	Medium      *float32        `firestore:"medium,omitempty" json:"medium,omitempty"`
	High        *float32        `firestore:"high,omitempty" json:"high,omitempty"`
	Huge        *float32        `firestore:"huge,omitempty" json:"huge,omitempty"`
	TooHigh     *float32        `firestore:"toohigh,omitempty" json:"toohigh,omitempty"`
	Snapshot    *gauge.Snapshot `firestore:"gauge,omitempty" json:"gauge,omitempty"`
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
	Measures    []Calibration `firestore:"measures,omitempty" json:"measures,omitempty"`
}

// CatalogueJSON is the JSON format to download the river catalogue
type CatalogueJSON struct {
	Version  string    `json:"version"`
	Sections []Section `json:"rivers"`
}

// RiverCache is an in-memory cache of sections
type RiverCache struct {
	URL     string
	Version string

	sectionMap map[string]Section
	rwMutex    *sync.RWMutex
}

// Load a section from the cache if it exists
func (c *RiverCache) Load(uuid string) (Section, bool) {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	s, ok := c.sectionMap[uuid]
	return s, ok
}

// Update polls for updated content version data
func (c *RiverCache) Update(ctx context.Context) (isChanged bool, s report.Span) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	span := report.StartSpan("rivers.update").Field("url", c.URL)

	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return false, span.End(err)
	}
	req = req.WithContext(ctx)

	client := &http.Client{
		Timeout: time.Second * 60,
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, span.End(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err := errors.New("Status code " + strconv.Itoa(resp.StatusCode))
		return false, span.End(err)
	}

	decoder := json.NewDecoder(resp.Body)
	data := CatalogueJSON{}
	if err := decoder.Decode(&data); err != nil {
		return false, span.End(err)
	}
	span = span.Field("version", data.Version)
	span = span.Field("stations_count", len(data.Sections))

	if c.Version == data.Version {
		return false, span.End()
	}

	c.rwMutex.Lock()
	{
		c.Version = data.Version
		c.sectionMap = make(map[string]Section, len(data.Sections))
		for _, s := range data.Sections {
			c.sectionMap[s.UUID] = s
		}
	}
	c.rwMutex.Unlock()

	return true, span.End()
}

// NewRiverCache creates a cache
func NewRiverCache(URL string) *RiverCache {
	return &RiverCache{
		URL:        URL,
		sectionMap: make(map[string]Section),
		rwMutex:    &sync.RWMutex{},
	}
}

// Each calls f sequentially for each section. If f returns false, each stops the iteration.
func (c *RiverCache) Each(f func(s Section) bool) {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

rangeLoop:
	for _, s := range c.sectionMap {
		if !f(s) {
			break rangeLoop
		}
	}
}

// Count returns the number of sections available
func (c *RiverCache) Count() int {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	return len(c.sectionMap)
}
