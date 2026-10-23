package main

import (
	"encoding/base64"
)

type Storer interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string)
	Keys() []string
	Len() int
}

func SetKeyWithEncryption(store Storer, key string, value string) (string, error) {
	var encoded = base64.StdEncoding.EncodeToString([]byte(value))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}

	return store.Get(key)
}
