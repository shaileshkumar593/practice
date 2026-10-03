package cache

import (
	"sync"
	"time"
)

type localItem struct {
	value     []byte
	expiresAt time.Time
}

type LocalCache struct {
	mu      sync.RWMutex
	items   map[string]localItem
	maxSize int
}

func NewLocalCache(maxSize int) *LocalCache {
	return &LocalCache{
		items:   make(map[string]localItem),
		maxSize: maxSize,
	}
}

func (c *LocalCache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	item, ok := c.items[key]

	if !ok {
		c.mu.RUnlock()
		return nil, false
	}

	if !time.Now().Before(item.expiresAt) {
		c.mu.RUnlock()
		c.Delete(key)
		return nil, false
	}

	value := append([]byte(nil), item.value...)
	c.mu.RUnlock()

	return value, true
}

func (c *LocalCache) Set(key string, value []byte, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.maxSize > 0 && len(c.items) >= c.maxSize {
		// Simple bounded eviction for the example.
		// Replace with a real LRU/TinyLFU cache in production.
		for k := range c.items {
			delete(c.items, k)
			break
		}
	}

	c.items[key] = localItem{
		value:     append([]byte(nil), value...),
		expiresAt: time.Now().Add(ttl),
	}
}

func (c *LocalCache) Delete(key string) {
	c.mu.Lock()
	delete(c.items, key)
	c.mu.Unlock()
}

func (c *LocalCache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}
