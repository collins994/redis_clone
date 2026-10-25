package main

import (
	"github.com/collins994/redis_clone/code/store"
	"github.com/collins994/redis_clone/code/store/kv"
)

func main() {
	var logger = NewLoggingMiddleware(kv.NewKVStore(0))
	logger.Set("name", "collins")
	logger.Get("name")

	store.SetKeyWithEncryption(logger, "name2", "collins ondeyo");
	logger.Get("name2")
	logger.Get("name3")
}
