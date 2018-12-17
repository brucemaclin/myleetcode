package lru

import (
	"container/list"
	"errors"
)

//LRUCache define a last recently used struct
type LRUCache struct {
	size      int
	evictList *list.List
	items     map[interface{}]*list.Element
	onEvict   EvictCallBack
}

//EvictCallBack when evict the lru call the function
type EvictCallBack func(key interface{}, val interface{})

type entry struct {
	key interface{}
	val interface{}
}

//NewLRUCache init a lru cache and return
//size should be positive
func NewLRUCache(size int, evict EvictCallBack) (*LRUCache, error) {
	if size <= 0 {
		return nil, errors.New("need positive size")
	}
	return &LRUCache{
		size:      size,
		evictList: list.New(),
		items:     make(map[interface{}]*list.Element),
		onEvict:   evict,
	}, nil
}

func (c *LRUCache) Add(key, val interface{}) (evicted bool) {
	if ent, ok := c.items[key]; ok {
		c.evictList.MoveToFront(ent)
		ent.Value.(*entry).val = val
		return false
	}
	ent := new(entry)
	ent.key = key
	ent.val = val
	entry := c.evictList.PushFront(ent)
	c.items[key] = entry
	length := c.evictList.Len()
	if length > c.size {
		c.removeOldest()
		return true
	}
	return false
}

func (c *LRUCache) Get(key interface{}) (val interface{}, ok bool) {
	if ent, ok := c.items[key]; ok {
		kv := ent.Value.(*entry)
		return kv.val, ok
	}
	return
}
func (c *LRUCache) removeOldest() {
	back := c.evictList.Back()
	if back != nil {
		c.removeElement(back)
	}
}
func (c *LRUCache) removeElement(ent *list.Element) {
	c.evictList.Remove(ent)
	kv := ent.Value.(*entry)
	delete(c.items, kv.key)
	if c.onEvict != nil {
		c.onEvict(kv.key, kv.val)
	}
}
