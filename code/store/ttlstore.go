package main

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrKeyExpired = errors.New("the key already expired")
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
			return ErrStoreCapacityReached
		}
	}
	if strings.TrimSpace(key) == "" {
		return ErrEmptyKey
	}

	s.data[key] = ttlEntry{
		value:      value,
		expiryTime: time.Now().Add(s.ttl),
	}

	return nil
}

func (s *TTLStore) Get(key string) (string, error) {
	if strings.TrimSpace(key) == "" {
		return "", ErrEmptyKey
	}

	val, ok := s.data[key]
	if !ok {
		return "", ErrKeyDoesNotExist
	}

	if time.Now().After(val.expiryTime) {
		s.Delete(key)
		return "", ErrKeyExpired
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
