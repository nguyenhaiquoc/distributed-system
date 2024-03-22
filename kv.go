package main

import (
	"errors"

	"github.com/rs/zerolog/log"
)

func init() {
	// Enable caller information globally
	log.Logger = log.With().Caller().Logger()
}

// KVStore represents a simple key-value store

var errKeyNotFound = errors.New("key not found")

type KVStore struct {
	data map[string]string
}

func NewKVStore() *KVStore {
	return &KVStore{data: make(map[string]string)}
}

func (kv *KVStore) Get(key string) (string, error) {
	value, ok := kv.data[key]

	// use zero log to debug the key and value received
	log.Debug().Str("key", key).Str("value", value).Msg("Get key and value")
	if !ok {
		return "", errKeyNotFound
	}
	return value, nil
}

func (kv *KVStore) Set(key string, value string) {
	kv.data[key] = value
}
