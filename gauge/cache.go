package gauge

import (
	"context"
	"sort"
	"sync"
	"time"
)

type Cache struct {
	snapMap   map[string]*CachedSnapshot
	rwMutex   *sync.RWMutex
	retention time.Duration
}

type CachedSnapshot struct {
	Station    Station
	Readings   []Reading
	ModifiedAt time.Time
}

type readingSorter []Reading

func (rs readingSorter) Len() int {
	return len(rs)
}

func (rs readingSorter) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs readingSorter) Less(i, j int) bool {
	return rs[i].DateTime.After(rs[j].DateTime)
}

func NewCache(ctx context.Context, retention time.Duration) *Cache {
	cache := Cache{
		rwMutex:   &sync.RWMutex{},
		retention: retention,
	}

	// spawn routine to regularly purge cache
	go func() {
		purgeInterval := time.Hour
		if cache.retention < purgeInterval {
			purgeInterval = cache.retention
		}

		ticker := time.NewTicker(purgeInterval)
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

func (c *Cache) Add(s Snapshot) {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	if cached, exists := c.snapMap[s.Station.UUID()]; !exists {
		c.snapMap[s.Station.DataURL] = &CachedSnapshot{
			Station:    s.Station,
			Readings:   s.Readings,
			ModifiedAt: time.Now(),
		}
	} else {
		cached.Readings = concat(cached.Readings, s.Readings)
		cached.ModifiedAt = time.Now()
	}
}

func (c *Cache) Get(uuid string) (CachedSnapshot, bool) {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()

	cached, exists := c.snapMap[uuid]
	return *cached, exists
}

func (c *Cache) purge() {
	c.rwMutex.Lock()
	defer c.rwMutex.Unlock()

	for k := range c.snapMap {
		removeOlderThan(time.Now().Add(-1*c.retention), &c.snapMap[k].Readings)
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
		if !found[x.DateTime] {
			found[x.DateTime] = true
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}

func removeOlderThan(keepSince time.Time, xs *[]Reading) {
	j := 0
	for i, x := range *xs {
		if keepSince.Before(x.DateTime) {
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}
