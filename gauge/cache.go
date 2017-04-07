package gauge

import (
	"context"
	"sort"
	"sync"
	"time"
)

type Cache struct {
	Map   map[string]*CachedSnapshot
	Mutex *sync.RWMutex
}

type CachedSnapshot struct {
	Station    Station
	Readings   []Reading
	ModifiedAt time.Time
}

type ReadingSorter []Reading

func (rs ReadingSorter) Len() int {
	return len(rs)
}

func (rs ReadingSorter) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
}

func (rs ReadingSorter) Less(i, j int) bool {
	return rs[i].DateTime.After(rs[j].DateTime)
}

func NewCache(ctx context.Context, duration time.Duration) *Cache {
	s := Cache{
		Mutex: &sync.RWMutex{},
	}

	// spawn routine to regularly purge cache
	go func() {
		purgeInterval := time.Hour
		if duration < purgeInterval {
			purgeInterval = duration
		}

		ticker := time.NewTicker(duration)
	purgeThenWait:
		for {
			s.Mutex.Lock()
			// DO PURGING
			s.Mutex.Unlock()

			select {
			case <-ticker.C:
			case <-ctx.Done():
				break purgeThenWait
			}
		}
		ticker.Stop()
	}()

	return &s
}

func (c *Cache) Add(s Snapshot) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if cached, exists := c.Map[s.Station.DataURL]; !exists {
		c.Map[s.Station.DataURL] = &CachedSnapshot{
			Station:    s.Station,
			Readings:   s.Readings,
			ModifiedAt: time.Now(),
		}
	} else {
		cached.Readings = concat(cached.Readings, s.Readings)
		cached.ModifiedAt = time.Now()
	}
}

func concat(a []Reading, b []Reading) []Reading {
	concat := append(a, b...)
	removeDuplicates(&concat)
	sort.Sort(ReadingSorter(concat))
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
		if keepSince.After(x.DateTime) {
			(*xs)[j] = (*xs)[i]
			j++
		}
	}
	*xs = (*xs)[:j]
}
