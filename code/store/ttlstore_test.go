package main

import (
	"errors"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	var store = NewTTLStore(time.Second*5, 1)

	t.Run("Set() with an empty key", func(t *testing.T) {
		if err := store.Set("", "Collins"); err == nil {
			t.Errorf("expected an error, got nil\n")
		}
	})

	t.Run("Set() before capacity is reached", func(t *testing.T) {
		if err := store.Set("name", "Collins"); err != nil {
			t.Errorf("%v\n", err)
		}
	})

	t.Run("Set() to update a key", func(t *testing.T) {
		if err := store.Set("name", "Fred"); err != nil {
			t.Errorf("Set() returned an error when updating a key value: %v\n", err)
		}
	})

	t.Run("Set() after capacity is reached", func(t *testing.T) {
		if err := store.Set("name2", "Fred"); err == nil || !errors.Is(err, ErrStoreCapacityReached) {
			t.Errorf("%v\n", err)
		}
	})

	t.Run("Set() after Delete and before capacity", func(t *testing.T) {
		store.Delete("name")
		if err := store.Set("name2", "Fred"); err != nil {
			t.Errorf("Set() returned an error %v\n", err)
		}
	})
}

func TestGet(t *testing.T) {
	var store = NewTTLStore(time.Second*2, 1)

	t.Run("Get() on an empty store", func(t *testing.T) {
		if _, err := store.Get("name"); err == nil || !errors.Is(err, ErrKeyDoesNotExist) {
			t.Errorf("Expected an ErrKeyDoesNotExist, got %v\n", err)
		}
	})

	store.Set("name", "Collins")
	t.Run("Get() with an empty key", func(t *testing.T) {
		if _, err := store.Get(" "); err == nil {
			t.Errorf("Expected an ErrEmptyKey, nil\n")
		}
	})

	time.Sleep(time.Second * 2)
	t.Run("Get() on an expired key", func(t *testing.T) {
		if _, err := store.Get("name"); err == nil || !errors.Is(err, ErrKeyExpired) {
			t.Errorf("Expected an ErrKeyExpired, got %v\n", err)
		}
	})

	t.Run("Get() on an evicted key", func(t *testing.T) {
		if _, err := store.Get("name"); err == nil || !errors.Is(err, ErrKeyDoesNotExist) {
			t.Errorf("Expected an ErrKeyDoesNotExist, got %v\n", err)
		}
	})
}
