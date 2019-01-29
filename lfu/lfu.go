package lfu

import (
	"container/list"
	"errors"
	"sync"
)

type LFUCache struct {
	size     int
	length   int
	freqList *list.List
	items    map[string]*cacheEntry
	lock     *sync.Mutex
}
type cacheEntry struct {
	key   string
	value interface{}
	node  *list.Element
}
type listEntry struct {
	freq  int
	items map[*cacheEntry]byte
}

func NewLFUCache(size int) (*LFUCache, error) {
	if size <= 0 {
		return nil, errors.New("size need to be positive")
	}
	return &LFUCache{
		size:     size,
		length:   0,
		freqList: list.New(),
		items:    make(map[string]*cacheEntry),
		lock:     new(sync.Mutex),
	}, nil
}

func (c *LFUCache) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if e, ok := c.items[key]; ok {
		e.value = value
		c.increment(e)
	} else {
		if c.length >= c.size {
			c.evict(1)
		}
		e = new(cacheEntry)
		e.key = key
		e.value = value
		c.items[key] = e
		c.increment(e)
		c.length++

	}
}
func (c *LFUCache) evict(count int) int {
	var evicted int
	entry := c.freqList.Front()
	for i := 0; i < count; {
		if entry == nil {
			return evicted
		}
		entries := entry.Value.(*listEntry).items
		for k := range entries {
			if i < count {
				i++
				c.length--
				delete(c.items, k.key)
				c.remEntry(entry, k)
				evicted++
			}
		}
	}
	return evicted
}

func (c *LFUCache) Get(key string) (val interface{}, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if e, ok := c.items[key]; ok {
		c.increment(e)
		return e.value, true
	} else {
		return nil, false
	}
}

func (c *LFUCache) increment(e *cacheEntry) {
	currentPlace := e.node
	var nextPlace *list.Element
	var nextFreq int
	if currentPlace == nil {
		nextFreq = 1
		nextPlace = c.freqList.Front()
	} else {
		nextFreq = currentPlace.Value.(*listEntry).freq + 1
		nextPlace = currentPlace.Next()
	}
	if nextPlace == nil || nextFreq != nextPlace.Value.(*listEntry).freq {
		li := new(listEntry)
		li.freq = nextFreq
		li.items = make(map[*cacheEntry]byte)
		if currentPlace != nil {
			nextPlace = c.freqList.InsertAfter(li, currentPlace)
		} else {
			nextPlace = c.freqList.PushFront(li)
		}
	}
	e.node = nextPlace
	e.node.Value.(*listEntry).items[e] = 1
	if currentPlace != nil {
		c.remEntry(currentPlace, e)
	}
}

func (c *LFUCache) remEntry(place *list.Element, entry *cacheEntry) {
	entries := place.Value.(*listEntry).items
	delete(entries, entry)
	if len(entries) == 0 {
		c.freqList.Remove(place)
	}
}
