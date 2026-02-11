package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu    sync.Mutex
	items map[string]cacheEntry
}

// Creates an in-memory cache.
func NewCache(interval time.Duration) *Cache {
	cache := &Cache{items: make(map[string]cacheEntry)}
	go cache.reapLoop(interval)
	return cache
}

// Add stores a cache entry with the provided key and value.
func (c *Cache) Add(k string, v []byte) {
	c.mu.Lock()
	c.items[k] = cacheEntry{
		createdAt: time.Now(),
		val:       v,
	}
	c.mu.Unlock()
}

// Get retrieves a cached value for the key, if present.
func (c *Cache) Get(k string) ([]byte, bool) {
	c.mu.Lock()
	data, ok := c.items[k]
	c.mu.Unlock()
	if !ok {
		return nil, false
	}
	return data.val, true
}

// reapLoop periodically removes expired cache entries.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		now := time.Now()
		c.mu.Lock()

		for key, entry := range c.items {
			timeLapsed := now.Sub(entry.createdAt)
			if timeLapsed >= interval {
				delete(c.items, key)
			}
		}

		c.mu.Unlock()
	}
}
