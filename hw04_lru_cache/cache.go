package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mx sync.RWMutex

	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) (hit bool) {
	var (
		li *ListItem
		ci cacheItem
	)
	c.mx.Lock()
	defer c.mx.Unlock()

	ci = cacheItem{
		key:   key,
		value: value,
	}

	if li, hit = c.items[key]; hit {
		li.Value = ci
		c.queue.MoveToFront(li)

		return
	}

	li = c.queue.PushFront(ci)
	c.items[key] = li

	if c.queue.Len() > c.capacity {
		last := c.queue.Back()
		c.queue.Remove(last)
		delete(c.items, last.Value.(cacheItem).key)
	}

	return
}

func (c *lruCache) Get(key Key) (value interface{}, hit bool) {
	var li *ListItem

	c.mx.RLock()
	defer c.mx.RUnlock()

	if li, hit = c.items[key]; !hit {
		return
	}

	value = li.Value.(cacheItem).value
	c.queue.MoveToFront(li)
	return
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
