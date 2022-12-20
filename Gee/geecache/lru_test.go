package geecache

import (
	lru2 "geecache/lru"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestLRUGet(t *testing.T) {
	lru := lru2.New(int64(0))
	lru.Add("key1", String("1234"))

	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed")

	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}
}
