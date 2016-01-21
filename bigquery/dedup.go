package main

import (
	"container/list"
)

// http://openmymind.net/Writing-An-LRU-Cache/

type DeDupeCache struct {
	capacity int
	items    map[string]*DeDupeCacheItem
	list     *list.List
}

type DeDupeCacheItem struct {
	key         string
	listElement *list.Element
}

func newDeDupeCache(capacity int) *DeDupeCache {
	return &DeDupeCache{
		capacity: capacity,
		items:    make(map[string]*DeDupeCacheItem, 10000),
		list:     list.New(),
	}
}

func (c *DeDupeCache) Exists(key string) bool {
	item, exists := c.items[key]
	if exists {
		c.promote(item)
	}
	return exists
}

func (c *DeDupeCache) promote(item *DeDupeCacheItem) {
	c.list.MoveToFront(item.listElement)
}

func (c *DeDupeCache) Set(key string) bool {
	if c.capacity < 1 {
		c.prune()
	}

	// still not enough room, fail
	if c.capacity < 1 {
		return false
	}

	item, exists := c.items[key]
	if exists {
		c.promote(item)
	} else {
		item = &DeDupeCacheItem{key: key}
		item.listElement = c.list.PushFront(item)
		c.items[key] = item
		c.capacity -= 1
	}
	return true
}

func (c *DeDupeCache) prune() {
	for i := 0; i < 50; i++ {
		tail := c.list.Back()
		if tail == nil {
			return
		}
		item := c.list.Remove(tail).(*DeDupeCacheItem)
		delete(c.items, item.key)
		c.capacity += 1
	}
}
