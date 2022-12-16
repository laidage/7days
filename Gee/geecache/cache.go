package geecache

import (
	"geecache/lru"
	"sync"
)

type cache struct {
	mu      sync.Mutex
	lru     *lru.Cache
	opacity int64
}

func (c *cache) get(key string) (value ByteView, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	v, ok := c.lru.Get(key)
	if ok {
		return v.(ByteView), true
	}
	return
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.opacity)
	}
	c.lru.Add(key, value)
}
