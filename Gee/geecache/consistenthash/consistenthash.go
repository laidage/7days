package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type HashFunc func([]byte) uint32

type CHash struct {
	replicas int
	hash     HashFunc
	hashMap  map[int]string
	keys     []int //sorted
}

func NewCHash(replicas int, fn HashFunc) *CHash {
	ch := &CHash{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if ch.hash == nil {
		ch.hash = crc32.ChecksumIEEE
	}
	return ch
}

func (ch *CHash) Add(machineNames ...string) {
	for _, machineName := range machineNames {
		for i := 0; i < ch.replicas; i++ {
			hashValue := int(ch.hash([]byte(strconv.Itoa(i) + machineName)))
			ch.keys = append(ch.keys, hashValue)
			ch.hashMap[hashValue] = machineName
		}
	}
	sort.Ints(ch.keys)
}

func (ch *CHash) Get(key string) string {
	if len(ch.keys) == 0 {
		return ""
	}
	hashValue := int(ch.hash([]byte(key)))
	idx := sort.Search(len(ch.keys), func(i int) bool {
		return ch.keys[i] >= hashValue
	})
	return ch.hashMap[ch.keys[idx%len(ch.keys)]]
}
