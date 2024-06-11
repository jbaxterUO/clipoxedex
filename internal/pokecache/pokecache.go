package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

type cacheEntry struct {
	value    []byte
	createAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		value:    value,
		createAt: time.Now(),
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, exists := c.cache[key]
	if !exists {
		return nil, false
	}
	return entry.value, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	intervalTimeAgo := time.Now().UTC().Add(-interval)
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createAt.Before(intervalTimeAgo) {
			delete(c.cache, k)
		}
	}
}
