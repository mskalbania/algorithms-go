package cache

import (
	"sync"
	"time"
)

type Cache struct {
	lock  sync.RWMutex
	ttl   time.Duration
	store map[string]*entry //consider pointer as val
}

type entry struct {
	value    string
	expireAt time.Time
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]*entry),
	}
}

func (c *Cache) Set(key, value string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.store[key] = &entry{
		value:    value,
		expireAt: time.Now().UTC().Add(c.ttl),
	}
}

func (c *Cache) Get(key string) string {
	c.lock.Lock()
	defer c.lock.Unlock()
	entry, ok := c.store[key]
	if ok {
		entry.expireAt = time.Now().UTC().Add(c.ttl)
		if time.Now().After(entry.expireAt) {
			delete(c.store, key)
			return ""
		}
		return entry.value
	}
	return ""
}
