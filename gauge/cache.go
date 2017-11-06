package gauge

import (
	"context"
	"math"
	"sort"
	"sync"
	"time"
)

// Cache is an in-memory cache of gauge stations and recent measurements
type Cache struct {
	snapMap          map[string]*Snapshot
	rwMutex          *sync.RWMutex
	defaultRetention time.Duration
	customRetention  map[string]time.Duration
}

// CacheStats is a collection of cache counts for monitoring telemetry
type CacheStats struct {
	StationCount         int
	CustomRetentionCount int
	AllReadingCount      int
	MaxReadingCount      int
	MinReadingCount      int
	OldestReading        time.Duration
}

type readingSorter []Reading

func (rs readingSorter) Len() int {
	return len(rs)
}

func (rs readingSorter) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs readingSorter) Less(i, j int) bool {
	return rs[i].EventTime.After(rs[j].EventTime)
}

// NewCache creates a new (empty) Cache
func NewCache(ctx context.Context, retention time.Duration) *Cache {
	cache := Cache{
		snapMap:          make(map[string]*Snapshot),
		rwMutex:          &sync.RWMutex{},
		defaultRetention: retention,
		customRetention:  make(map[string]time.Duration),
	}

	// spawn routine to regularly purge cache
	go func() {
		interval := time.Hour
		if cache.defaultRetention < interval {
			interval = cache.defaultRetention
		}

		ticker := time.NewTicker(interval)
	purgeThenWait:
		for {
			cache.purge()

			select {
			case <-ticker.C:
			case <-ctx.Done():
				break purgeThenWait
			}
		}
		ticker.Stop()
	}()

	return &cache
}

// ChangeRetention changes the retention period for a particular station
func (c *Cache) ChangeRetention(dataURL string, retention time.Duration) {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	c.customRetention[dataURL] = retention
	return
}

// retentionInLock is a helper method to get configured retention once in a locked state
func (c *Cache) retentionInLock(dataURL string) time.Duration {
	r, ok := c.customRetention[dataURL]
	if !ok {
		return c.defaultRetention
	}
	return r
}

// Add includes the provided Snapshot in the cached dataset
func (c *Cache) Add(s *Snapshot) {
	key := s.Station.DataURL

	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	retention := c.retentionInLock(s.Station.DataURL)
	removeOlderThan(time.Now().Add(-1*retention), &s.Readings)
	item, exists := c.snapMap[key]
	if !exists {
		item = &Snapshot{
			Station: s.Station,
		}
		c.snapMap[key] = item
	}
	item.Readings = concat(item.Readings, s.Readings)
	item.ProcessingTime = time.Now()
}

// Load retrieves the cached Snapshot of a particular station if available
func (c *Cache) Load(dataURL string) (Snapshot, bool) {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	cached, exists := c.snapMap[dataURL]
	if exists {
		return *cached, exists
	}
	return Snapshot{}, false
}

// Count returns the number of stations
func (c *Cache) Count() int {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	return len(c.snapMap)
}

// Each calls f sequentially for each snapshot. If f returns false, each stops the iteration.
func (c *Cache) Each(f func(s *Snapshot) bool) {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

rangeLoop:
	for _, s := range c.snapMap {
		if !f(s) {
			break rangeLoop
		}
	}
}

// Stats returns a collection of cache counts for monitoring telemetry
func (c *Cache) Stats() CacheStats {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	var oldest time.Time
	all := 0
	max := 0
	min := math.MaxInt64
	for k := range c.snapMap {
		len := len(c.snapMap[k].Readings)
		all += len
		if len > max {
			max = len
		}
		if len < min {
			min = len
		}
		if len > 0 {
			last := c.snapMap[k].Readings[len-1]
			if oldest.IsZero() || last.EventTime.Before(oldest) {
				oldest = last.EventTime
			}
		}
	}
	if min == math.MaxInt64 {
		min = 0
	}
	status := CacheStats{
		StationCount:         len(c.snapMap),
		CustomRetentionCount: len(c.customRetention),
		AllReadingCount:      all,
		MaxReadingCount:      max,
		MinReadingCount:      min,
	}
	if !oldest.IsZero() {
		status.OldestReading = time.Now().Sub(oldest)
	}

	return status
}

func (c *Cache) purge() {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	for k := range c.snapMap {
		r := c.retentionInLock(c.snapMap[k].Station.DataURL)
		removeOlderThan(time.Now().Add(-1*r), &c.snapMap[k].Readings)
	}
}

func concat(a []Reading, b []Reading) []Reading {
	concat := append(a, b...)
	removeDuplicates(&concat)
	sort.Sort(readingSorter(concat))
	return concat
}

func removeDuplicates(xs *[]Reading) {
	found := make(map[time.Time]bool)
	j := 0
	for i, x := range *xs {
		if !found[x.EventTime] {
			found[x.EventTime] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

func removeOlderThan(keepSince time.Time, xs *[]Reading) {
	j := 0
	for i, x := range *xs {
		if keepSince.Before(x.EventTime) {
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}
