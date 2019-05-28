package hashmap

import (
	"crypto/md5"
	"fmt"
)

type Item struct {
	Key   string
	Value string
}

type HashMap struct {
	length int64
	bucket map[string][]Item
}

func New() *HashMap {
	return &HashMap{length: 0, bucket: make(map[string][]Item)}
}

func (m *HashMap) Set(key, value string) {
	var (
		ok    bool
		items []Item
		hash  = hashfunc(key)
	)
	if items, ok = m.bucket[hash]; ok {
		for _, item := range items {
			if item.Key == key {
				item.Key = value
				return
			}
		}
	}
	m.bucket[hash] = append(items, Item{Key: key, Value: value})
	m.length++
}

func (m *HashMap) Get(key string) (string, bool) {
	var hash = hashfunc(key)
	if items, ok := m.bucket[hash]; ok {
		for _, item := range items {
			if item.Key == key {
				return item.Value, true
			}
		}
	}
	return "", false
}

func (m *HashMap) Remove(key string) {
	var hash = hashfunc(key)
	if items, ok := m.bucket[hash]; ok {
		for i := 0; i < len(items); i++ {
			if items[i].Key == key {
				m.bucket[hash] = append(items[:i], items[i+1:]...)
				m.length--
				return
			}
		}
	}
}

func (m *HashMap) Len() int64 {
	return m.length
}

func hashfunc(value string) string {
	return fmt.Sprintf("%x\n", md5.Sum([]byte(value)))
}
