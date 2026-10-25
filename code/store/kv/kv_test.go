package kv

import (
	"errors"
	"reflect"
	"testing"
	"github.com/collins994/redis_clone/code/store"
)

func TestKeys_ReturnsSorted(t *testing.T) {
	kv := NewKVStore(0)
	kv.Set("1", "collins")
	kv.Set("3", "Fred")
	kv.Set("2", "Steve")

	want := []string{"1", "2", "3"}
	got := kv.Keys()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want = %v, got = %v\n", want, got)
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	kv := NewKVStore(0)
	kv.Set("name", "collins")

	_, err := kv.Get("name")
	if err != nil {
		t.Errorf("%v\n", err)
	}

	kv.Delete("name")
	_, err = kv.Get("name")
	if err == nil {
		t.Errorf("expected %v, got nil\n", err)
	}
}

func TestDelete(t *testing.T) {
	var tests = []struct {
		name      string
		setup     map[string]string
		deleteKey string
		wantLen   int
	}{
		{"deletes an existing key", map[string]string{"name": "collins", "name2": "Fred"}, "name", 1},
		{"no op on missing key", map[string]string{"name": "collins"}, "name2", 1},
		{"no op on empty kv", map[string]string{}, "name2", 0},
		{"deletes last remaining key", map[string]string{"name": "collins"}, "name", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var kv = NewKVStore(0)
			for key, val := range test.setup {
				kv.Set(key, val)
			}

			kv.Delete(test.deleteKey)

			if gotLen := kv.Len(); gotLen != test.wantLen {
				t.Errorf("after Delete(%v), want len=%v, got len=%v\n", test.deleteKey, gotLen, test.wantLen)
			}

			if _, err := kv.Get(test.deleteKey); err == nil {
				t.Errorf("after Delete(%v), key still present", test.deleteKey)
			}
		})
	}
}

func TestGetSet_EmptyKeys(t *testing.T) {
	var kv = NewKVStore(0)

	if _, err := kv.Get(" "); err == nil || !errors.Is(err, store.ErrEmptyKey) {
		t.Errorf("Get() failed to return an error on an empty key\n")
	}

	if err := kv.Set(" ", ""); err == nil || !errors.Is(err, store.ErrEmptyKey) {
		t.Errorf("Set() failed to return an error on an empty key\n")
	}
}

func TestSet_ReturnsAnErrorAtStoreCapacity(t *testing.T) {
	kv := NewKVStore(1)
	kv.Set("name", "collins")

	if err := kv.Set("name3", "Fred"); err == nil || !errors.Is(err, store.ErrStoreCapacityReached) {
		t.Errorf("Set() failed to return an error at max store capacity\n")
	}

	// updating a key value
	if err := kv.Set("name", "Fred"); err != nil {
		t.Errorf("store.Set() returned an error when updating a key already in the store\n")
	}

	kv.Delete("name")
	if err := kv.Set("name3", "Fred"); err != nil {
		t.Errorf("Set() returned an error yet store capacity is not reached\n")
	}
}
