package main

import (
	"github.com/collins994/redis_clone/code/store"
	"log"
	"os"
)

type LoggingMiddleware struct {
	inner  store.Storer
	logger *log.Logger
}

func NewLoggingMiddleware(inner store.Storer) store.Storer {
	return &LoggingMiddleware{
		inner:  inner,
		logger: log.New(os.Stdout, "[Log ]", log.Ltime|log.Lmicroseconds),
	}
}

func (l *LoggingMiddleware) Get(key string) (string, error) {
	val, err := l.inner.Get(key);
	if err != nil {
		l.logger.Printf("Get %v -> miss(%v)\n", key, err);
	} else {
		l.logger.Printf("Get %q -> hit(%v)\n", key, val);
	}
	return val, err;
}

func (l *LoggingMiddleware) Set(key, value string) error {
	l.logger.Printf("Set %v\n", key);
	return l.inner.Set(key,value);
}

func (l *LoggingMiddleware) Delete(key string) {
	l.logger.Printf("Delete %v\n", key);
	l.inner.Delete(key);
}

func (l *LoggingMiddleware) Keys() []string {
	var gotKeys = l.inner.Keys();
	l.logger.Printf("Keys %v\n", gotKeys);
	return gotKeys;
}

func (l *LoggingMiddleware) Len() int {
	var gotLen = l.inner.Len();
	l.logger.Printf("Len %v\n", gotLen);
	return gotLen;
}
