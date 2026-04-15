package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mux     *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		entries: make(map[string]cacheEntry),
		mux: &sync.Mutex,
	}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.enc.entries[key] = caccacheEntry{
		createdAt;  time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mux.Lock()
	defer c.mux.Unlock()

	if c.entries[key] {
		return c.entries[key].val, true
	}
	return nil, false
}

func (c *Cache) reapLoop(interval time.Time) {

	c.mux.Lock()
	defer c.mux.Unlock()

	ticker := time.NewTicker(interval)

	for range ticker.C{
		for i, entrie := range c.entries {
			if time.Now() - entrie.createdAt > interval {
				delete(c.entries, i)
			}
	}
}
}

