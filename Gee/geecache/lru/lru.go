package lru

import "C"
import (
	"container/list"
)

type Cache struct {
	ll       *list.List
	cache    map[string]*list.Element
	maxBytes int64
	nBytes   int64
}

type Entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64) *Cache {
	return &Cache{
		cache:    make(map[string]*list.Element),
		maxBytes: maxBytes,
		nBytes:   0,
		ll:       list.New(),
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*Entry)
		return kv.value, true
	}
	return
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*Entry)
		c.nBytes += int64(value.Len() - kv.value.Len())
		kv.value = value
	} else {
		entry := &Entry{
			key:   key,
			value: value,
		}
		ele := c.ll.PushFront(entry)
		c.cache[key] = ele
		c.nBytes += int64(value.Len() + len(key))
	}
	for c.maxBytes != 0 && c.maxBytes < c.nBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		kv := ele.Value.(*Entry)
		delete(c.cache, kv.key)
		c.ll.Remove(ele)
		c.nBytes -= int64(len(kv.key) + kv.value.Len())
	}
}
