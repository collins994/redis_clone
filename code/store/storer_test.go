package store

import (
	"testing"
	"time"
	"github.com/collins994/redis_clone/code/store/ttlstore"
	"github.com/collins994/redis_clone/code/store/kv"
)

func TestSetKeyWithEncryption_kv(t *testing.T) {
	var kv = NewKVStore(0);
	t.Run("SetKeyWithEncryption() on kv store", func(t *testing.T) {
		if _, err := SetKeyWithEncryption(kv, "name", "Collins"); err != nil {
			t.Errorf("Expected nil, got %v\n", err)
		}
	})

	t.Run("Get() after SetKeyWithEncryption()", func(t *testing.T){
		if _, err := kv.Get("name"); err != nil {
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
