package geecache

import (
	"fmt"
	"log"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	groups = make(map[string]*Group, 0)
	mu     sync.RWMutex
)

type Group struct {
	name       string
	mainCache  cache
	getter     GetterFunc
	peerPicker PeerPicker
}

func NewGroup(name string, opacity int64, getter GetterFunc) *Group {
	if getter == nil {
		panic("err getter")
	}
	mu.Lock()
	defer mu.Unlock()
	group := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{opacity: opacity},
	}
	groups[name] = group
	return group
}

func GetGroup(name string) *Group {
	mu.RLock()
	group := groups[name]
	mu.RUnlock()
	return group
}
func (group *Group) RegisterPicker(peerPicker PeerPicker) {
	if group.peerPicker != nil {
		panic("peerPicker could not be more than one.")
	}
	group.peerPicker = peerPicker
}

func (group *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is required")
	}
	value, ok := group.mainCache.get(key)
	if ok {
		log.Println("[GeeCache] hit")
		return value, nil
	}
	return group.load(key)
}

func (group *Group) load(key string) (value ByteView, err error) {
	if group.peerPicker != nil {
		if getter, ok := group.peerPicker.Pick(key); ok {
			if value, err = group.getFromPeer(getter, key); err == nil {
				return value, nil
			}
			log.Println("[GeeCache] Failed to get from peer", err)
		}
	}
	return group.getLocally(key)
}

func (group *Group) getFromPeer(getter PeerGetter, key string) (ByteView, error) {
	bytes, err := getter.Get(group.name, key)
	if err != nil {
		return ByteView{}, err
	}
	return ByteView{b: bytes}, nil
}

func (group *Group) getLocally(key string) (ByteView, error) {
	bytes, err := group.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b: bytes}
	group.mainCache.add(key, value)
	return value, nil
}
