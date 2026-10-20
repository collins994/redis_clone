package main

import (
	// "fmt"
	"errors"
	"sort"
	"strings"
)

var ErrKeyDoesNotExist = errors.New("the key does not exist")
var ErrEmptyKey = errors.New("the key is empty")
var ErrStoreCapacityReached = errors.New("The Store capacity has been reached")

type Store struct {
	data map[string]string
	capacity int;
}

func NewStore(capacity int) *Store {
	return &Store{
		data: make(map[string]string),
		capacity: capacity,
	}
}

func (s *Store) Set(key, value string) error {
	if s.Len() >= s.capacity {
		if _, exists := s.data[key]; !exists{ // a new key at store capacity
			return ErrStoreCapacityReached;
		}
	}

	if strings.TrimSpace(key) == "" {
		return ErrEmptyKey
	}

	s.data[key] = value
	return nil
}

func (s *Store) Get(key string) (string, error) {
	if strings.TrimSpace(key) == "" {
		return "", ErrEmptyKey
	}

	val, ok := s.data[key]
	if !ok {
		return "", ErrKeyDoesNotExist
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
