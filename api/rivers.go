package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/rainchasers/report"
)

//  Catalogue from https://app.rainchasers.com/catalogue.json like:
//  { version: "v138",
//    rivers: [ {
//	    uuid: "14be0011-a293-4e0d-89df-c0216cf9fe5e",
//	    river: "Aeron",
//	    section: "Cilau Aeron to Aberaeron",
//	    km: 9,
//	    grade: {"text":"2/3","value":2.5,"max":null},
//	    putin: {"lat":52.2151638,"lng":-4.197429},
//	    takeout: {"lat":52.2436576,"lng":-4.2644627},
//	    desc: "Good training run ...",
//	    directions: "The takeout is at ..."
//    }, { ... etc ... } ]
//  }

// LatLng is a geographic location
type LatLng struct {
	Lat float32 `json:"lat"`
	Lng float32 `json:"lng"`
}

// Grade is the section difficultly
type Grade struct {
	Human   string   `json:"text"`
	Average float32  `json:"value"`
	Max     *float32 `json:"max"`
}

// Calibration is a referenced gauge related to a section
type Calibration struct {
	Reference   string   `json:"ref"`
	Description string   `json:"desc"`
	Scrape      *float64 `json:"scrape"`
	Low         *float64 `json:"low"`
	Medium      *float64 `json:"medium"`
	High        *float64 `json:"high"`
	Huge        *float64 `json:"huge"`
	TooHigh     *float64 `json:"toohigh"`
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
	Gauges      []Calibration `json:"gauge"`
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

	log        *report.Logger
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
func (c *RiverCache) Update() error {
	tick := c.log.Tick()
	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	req, err := http.NewRequest("GET", c.URL, nil)
	if err != nil {
		<-c.log.Action("river.update.failed", report.Data{
			"url":   c.URL,
			"error": err.Error(),
			"step":  "setup",
		})
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		<-c.log.Action("river.update.failed", report.Data{
			"url":   c.URL,
			"error": err.Error(),
			"step":  "request",
		})
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		<-c.log.Action("river.update.failed", report.Data{
			"url":   c.URL,
			"error": "Status code " + strconv.Itoa(resp.StatusCode),
			"step":  "request",
		})
		return errors.New("Status code " + strconv.Itoa(resp.StatusCode))
	}
	c.log.Tock(tick, "river.update.downloaded", report.Data{
		"url": c.URL,
		"len": resp.ContentLength,
	})

	decoder := json.NewDecoder(resp.Body)
	data := CatalogueJSON{}
	if err := decoder.Decode(&data); err != nil {
		<-c.log.Action("river.update.failed", report.Data{
			"url":   c.URL,
			"error": err.Error(),
			"step":  "decode",
		})
		return err
	}

	if c.Version == data.Version {
		c.log.Tock(tick, "river.update.unchanged", report.Data{"url": c.URL})
		return nil
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
	c.log.Tock(tick, "river.update.changed", report.Data{
		"url":     c.URL,
		"version": data.Version,
		"count":   len(data.Sections),
	})
	return nil
}

// NewRiverCache creates a cache and populates it
func NewRiverCache(ctx context.Context, URL string, log *report.Logger) (*RiverCache, error) {
	cache := &RiverCache{
		URL:        URL,
		log:        log,
		sectionMap: make(map[string]Section),
		rwMutex:    &sync.RWMutex{},
	}

	if err := cache.Update(); err != nil {
		return cache, err
	}

	// spawn routine to poll for content updates
	go func() {
		ticker := time.NewTicker(60 * time.Second)
	updateThenWait:
		for {
			select {
			case <-ticker.C:
			case <-ctx.Done():
				break updateThenWait
			}

			cache.Update()
		}
		ticker.Stop()
	}()

	return cache, nil
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
