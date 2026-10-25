package kv

import (
	"sort"
	"strings"
	"github.com/collins994/redis_clone/code/store"
)

type Store struct {
	data     map[string]string
	capacity int
}

func NewKVStore(capacity int) *Store {
	return &Store{
		data:     make(map[string]string),
		capacity: capacity,
	}
}

func (s *Store) Set(key, value string) error {
	if s.Len() >= s.capacity && s.capacity != 0 {
		if _, exists := s.data[key]; !exists { // a new key at store capacity
			return store.ErrStoreCapacityReached
		}
	}

	if strings.TrimSpace(key) == "" {
		return store.ErrEmptyKey
	}

	s.data[key] = value
	return nil
}

func (s *Store) Get(key string) (string, error) {
	if strings.TrimSpace(key) == "" {
		return "", store.ErrEmptyKey
	}

	val, ok := s.data[key]
	if !ok {
		return "", store.ErrKeyDoesNotExist
	}
	return val, nil
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for key, _ := range s.data {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

func (s *Store) Len() int {
	return len(s.data)
}
