package main

import (
	"testing"
)

func TestKeyValueStore(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		existent bool
		expected string
	}{
		{
			name:     "Test SET operation",
			key:      "key1",
			value:    "value1",
			existent: true,
			expected: "value1",
		},
		{
			name:     "Test GET operation for a non-existent key",
			key:      "key2",
			value:    "",
			existent: false,
			expected: "",
		},
	}

	// Create a new instance of the key-value store, initializing data
	kv := NewKVStore()
	kv.Set("key1", "value1")
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test SET operation
			value, err := kv.Get(test.key)
			if err != nil {
				if (test.existent && err != errKeyNotFound) || (!test.existent && err == nil) {
					t.Errorf("Error getting value for %s: %v", test.key, err)
				}
			} else {
				if value != test.expected {
					t.Errorf("Expected %s, got %s", test.expected, value)
				}
			}
		})
	}
}
