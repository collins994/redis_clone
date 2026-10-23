package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestKeys_ReturnsSorted(t *testing.T) {
	store := NewStore(0)
	store.Set("1", "collins")
	store.Set("3", "Fred")
	store.Set("2", "Steve")

	want := []string{"1", "2", "3"}
	got := store.Keys()

	if !reflect.DeepEqual(want, got) {
		t.Errorf("want = %v, got = %v\n", want, got)
	}
}

func TestSetGet_RoundTrip(t *testing.T) {
	store := NewStore(0)
	store.Set("name", "collins")

	_, err := store.Get("name")
	if err != nil {
		t.Errorf("%v\n", err)
	}

	store.Delete("name")
	_, err = store.Get("name")
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
		{"no op on empty store", map[string]string{}, "name2", 0},
		{"deletes last remaining key", map[string]string{"name": "collins"}, "name", 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var store = NewStore(0)
			for key, val := range test.setup {
				store.Set(key, val)
			}

			store.Delete(test.deleteKey)

			if gotLen := store.Len(); gotLen != test.wantLen {
				t.Errorf("after Delete(%v), want len=%v, got len=%v\n", test.deleteKey, gotLen, test.wantLen)
			}

			if _, err := store.Get(test.deleteKey); err == nil {
				t.Errorf("after Delete(%v), key still present", test.deleteKey)
			}
		})
	}
}

func TestGetSet_EmptyKeys(t *testing.T) {
	var store = NewStore(0)

	if _, err := store.Get(" "); err == nil || !errors.Is(err, ErrEmptyKey) {
		t.Errorf("Get() failed to return an error on an empty key\n")
	}

	if err := store.Set(" ", ""); err == nil || !errors.Is(err, ErrEmptyKey) {
		t.Errorf("Set() failed to return an error on an empty key\n")
	}
}

func TestSet_ReturnsAnErrorAtStoreCapacity(t *testing.T) {
	store := NewStore(1)
	store.Set("name", "collins")

	if err := store.Set("name3", "Fred"); err == nil || !errors.Is(err, ErrStoreCapacityReached) {
		t.Errorf("Set() failed to return an error at max store capacity\n")
	}

	// updating a key value
	if err := store.Set("name", "Fred"); err != nil {
		t.Errorf("store.Set() returned an error when updating a key already in the store\n")
	}

	store.Delete("name")
	if err := store.Set("name3", "Fred"); err != nil {
		t.Errorf("Set() returned an error yet store capacity is not reached\n")
	}
}
