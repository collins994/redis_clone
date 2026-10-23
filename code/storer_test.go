package main

import (
	"testing"
	"time"
	"github.com/collins994/redis_clone/code/store"
)

func TestSetKeyWithEncryption_store(t *testing.T) {
	var store = NewStore(0);
	t.Run("SetKeyWithEncryption() on store", func(t *testing.T) {
		if _, err := SetKeyWithEncryption(store, "name", "Collins"); err != nil {
			t.Errorf("Expected nil, got %v\n", err)
		}
	})

	t.Run("Get() after SetKeyWithEncryption()", func(t *testing.T){
		if _, err := store.Get("name"); err != nil {
			t.Errorf("Expected nil, got %v\n", err)
		}
	})
}

func TestSetKeyWithEncryption_TTLStore(t *testing.T) {
	var store = NewTTLStore(time.Second*5, 0);
	t.Run("SetKeyWithEncryption() on ttlstore", func(t *testing.T) {
		if _, err := SetKeyWithEncryption(store, "name", "Collins"); err != nil {
			t.Errorf("Expected nil, got %v\n", err)
		}
	})

	t.Run("Get() after SetKeyWithEncryption()", func(t *testing.T){
		if _, err := store.Get("name"); err != nil {
			t.Errorf("Expected nil, got %v\n", err)
		}
	})
}
