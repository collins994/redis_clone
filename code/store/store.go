package store

import (
	"errors"
)

var ErrKeyDoesNotExist = errors.New("the key does not exist")
var ErrEmptyKey = errors.New("the key is empty")
var ErrStoreCapacityReached = errors.New("The Store capacity has been reached")
var ErrKeyExpired = errors.New("the key already expired")
