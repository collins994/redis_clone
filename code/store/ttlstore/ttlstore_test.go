package ttlstore

import (
	"errors"
	"testing"
	"time"
	"github.com/collins994/redis_clone/code/store"
)

func TestSet(t *testing.T) {
	var ttlStore = NewTTLStore(time.Second*5, 1)

	t.Run("Set() with an empty key", func(t *testing.T) {
		if err := ttlStore.Set("", "Collins"); err == nil {
			t.Errorf("expected an error, got nil\n")
		}
	})

	t.Run("Set() before capacity is reached", func(t *testing.T) {
		if err := ttlStore.Set("name", "Collins"); err != nil {
			t.Errorf("%v\n", err)
		}
	})

	t.Run("Set() to update a key", func(t *testing.T) {
		if err := ttlStore.Set("name", "Fred"); err != nil {
			t.Errorf("Set() returned an error when updating a key value: %v\n", err)
		}
	})

	t.Run("Set() after capacity is reached", func(t *testing.T) {
		if err := ttlStore.Set("name2", "Fred"); err == nil || !errors.Is(err, store.ErrStoreCapacityReached) {
			t.Errorf("%v\n", err)
		}
	})

	t.Run("Set() after Delete and before capacity", func(t *testing.T) {
		ttlStore.Delete("name")
		if err := ttlStore.Set("name2", "Fred"); err != nil {
			t.Errorf("Set() returned an error %v\n", err)
		}
	})
}

func TestGet(t *testing.T) {
	var ttlStore = NewTTLStore(time.Second*2, 1)

	t.Run("Get() on an empty store", func(t *testing.T) {
		if _, err := ttlStore.Get("name"); err == nil || !errors.Is(err, store.ErrKeyDoesNotExist) {
			t.Errorf("Expected an ErrKeyDoesNotExist, got %v\n", err)
		}
	})

	ttlStore.Set("name", "Collins")
	t.Run("Get() with an empty key", func(t *testing.T) {
		if _, err := ttlStore.Get(" "); err == nil {
			t.Errorf("Expected an ErrEmptyKey, nil\n")
		}
	})

	time.Sleep(time.Second * 2)
	t.Run("Get() on an expired key", func(t *testing.T) {
		if _, err := ttlStore.Get("name"); err == nil || !errors.Is(err, store.ErrKeyExpired) {
			t.Errorf("Expected an ErrKeyExpired, got %v\n", err)
		}
	})

	t.Run("Get() on an evicted key", func(t *testing.T) {
		if _, err := ttlStore.Get("name"); err == nil || !errors.Is(err, store.ErrKeyDoesNotExist) {
			t.Errorf("Expected an ErrKeyDoesNotExist, got %v\n", err)
		}
	})
}
