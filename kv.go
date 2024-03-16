package main

// KVStore represents a simple key-value store

type KVStore struct {
	data map[string]string
}

func NewKVStore() *KVStore {
	return &KVStore{data: make(map[string]string)}
}

func (kv *KVStore) Get(key string) (string, error) {
	return kv.data[key], nil
}

func (kv *KVStore) Set(key string, value string) {
	kv.data[key] = value
}
