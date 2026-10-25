package main

import (
	"github.com/collins994/redis_clone/code/store"
	"github.com/collins994/redis_clone/code/store/kv"
)

func main() {
	var kvStore = NewMetricsMiddleware(NewLoggingMiddleware(kv.NewKVStore(0)))
	kvStore.Set("name", "collins")
	kvStore.Get("name")

	store.SetKeyWithEncryption(kvStore, "name2", "collins ondeyo");
	kvStore.Get("name2")
	kvStore.Get("name3")
	kvStore.Keys()

	kvStore.Report();
}
