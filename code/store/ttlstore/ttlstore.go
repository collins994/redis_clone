package ttlstore

import (
	"github.com/collins994/redis_clone/code/store"
	"strings"
	"time"
)

type ttlEntry struct {
	value      string
	expiryTime time.Time
}

type TTLStore struct {
	data     map[string]ttlEntry
	ttl      time.Duration
	capacity int
}

func NewTTLStore(ttl time.Duration, capacity int) *TTLStore {
	return &TTLStore{
		data:     make(map[string]ttlEntry),
		ttl:      ttl,
		capacity: capacity,
	}
}

func (s *TTLStore) Set(key string, value string) error {
	if s.Len() >= s.capacity && s.capacity != 0 {
		if _, exists := s.data[key]; !exists { // a new key at store capacity
			return store.ErrStoreCapacityReached
		}
	}
	if strings.TrimSpace(key) == "" {
		return store.ErrEmptyKey
	}

	s.data[key] = ttlEntry{
		value:      value,
		expiryTime: time.Now().Add(s.ttl),
	}

	return nil
}

func (s *TTLStore) Get(key string) (string, error) {
	if strings.TrimSpace(key) == "" {
		return "", store.ErrEmptyKey
	}

	val, ok := s.data[key]
	if !ok {
		return "", store.ErrKeyDoesNotExist
	}

	if time.Now().After(val.expiryTime) {
		s.Delete(key)
		return "", store.ErrKeyExpired
	}
	return val.value, nil
}

func (s *TTLStore) Delete(key string) {
	delete(s.data, key)
}

func (s *TTLStore) Len() int {
	return len(s.data)
}

func (s *TTLStore) Keys() []string {
	var keys = []string{}
	for key, _ := range s.data {
		keys = append(keys, key)
	}
	return keys
}
