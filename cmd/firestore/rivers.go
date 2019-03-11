package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/rainchasers/com.rainchasers.gauge/internal/daemon"
	"github.com/rainchasers/com.rainchasers.gauge/internal/gauge"
	"github.com/rainchasers/report"
)

//  Catalogue from https://app.rainchasers.com/catalogue.json

// LatLng is a geographic location
type LatLng struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

// Grade is the section difficultly
type Grade struct {
	Human   string   `json:"text"`
	Average float32  `json:"value"`
	Max     *float32 `json:"max,omitempty"`
}

// Calibration is a referenced gauge related to a section
type Calibration struct {
	DataURL     string          `json:"data_url"`
	Description string          `json:"desc"`
	Scrape      *float32        `json:"scrape,omitempty"`
	Low         *float32        `json:"low,omitempty"`
	Medium      *float32        `json:"medium,omitempty"`
	High        *float32        `json:"high,omitempty"`
	Huge        *float32        `json:"huge,omitempty"`
	TooHigh     *float32        `json:"toohigh,omitempty"`
	Snapshot    *gauge.Snapshot `json:"gauge,omitempty"`
}

// Section is a paddleable river section
type Section struct {
	UUID        string        `json:"uuid"`
	SectionName string        `json:"section"`
	RiverName   string        `json:"river"`
	KM          float32       `json:"km"`
	Grade       Grade         `json:"grade"`
	Putin       LatLng        `json:"putin"`
	Takeout     LatLng        `json:"takeout"`
	Description string        `json:"desc"`
	Directions  string        `json:"directions"`
	Measures    []Calibration `json:"measures,omitempty"`
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
func (c *RiverCache) Update(ctx context.Context, d *daemon.Supervisor) (isChanged bool, err error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	ctx = d.StartSpan(ctx, "rivers.update")
	defer func() {
		d.EndSpan(ctx, err, report.Data{
			"url":     c.URL,
			"version": c.Version,
			"count":   c.Count(),
		})
		cancel()
	}()

	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		return false, err
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false, errors.New("Status code " + strconv.Itoa(resp.StatusCode))
	}

	decoder := json.NewDecoder(resp.Body)
	data := CatalogueJSON{}
	if err := decoder.Decode(&data); err != nil {
		return false, err
	}

	if c.Version == data.Version {
		return false, nil
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
	return true, nil
}

// NewRiverCache creates a cache and populates it
func NewRiverCache(ctx context.Context, d *daemon.Supervisor, URL string) (*RiverCache, error) {
	cache := &RiverCache{
		URL:        URL,
		sectionMap: make(map[string]Section),
		rwMutex:    &sync.RWMutex{},
	}

	_, err := cache.Update(ctx, d)
	return cache, err
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
