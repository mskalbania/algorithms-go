package other

import (
	"container/list"
	"sync"
)

type LRUCache[K comparable, V any] struct {
	capacity   int
	cache      map[K]*list.Element
	accessList *list.List
	lock       sync.Mutex
}

type entry[K comparable, V any] struct {
	key K
	val V
}

func NewLRUCache[K comparable, V any](capacity int) *LRUCache[K, V] {
	return &LRUCache[K, V]{
		capacity:   capacity,
		cache:      make(map[K]*list.Element),
		accessList: list.New(),
	}
}

func (c *LRUCache[K, V]) Put(key K, val V) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if element, ok := c.cache[key]; ok {
		c.accessList.MoveToFront(element)
		element.Value.(*entry[K, V]).val = val
		return
	}

	if c.accessList.Len() == c.capacity {
		back := c.accessList.Back()
		c.accessList.Remove(back)
		delete(c.cache, back.Value.(*entry[K, V]).key)
		return
	}

	newElement := c.accessList.PushFront(&entry[K, V]{key, val})
	c.cache[key] = newElement
}

func (c *LRUCache[K, V]) Get(key K) (V, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if element, ok := c.cache[key]; ok {
		c.accessList.MoveToFront(element)
		return element.Value.(*entry[K, V]).val, true
	} else {
		var zero V
		return zero, false
	}
}
