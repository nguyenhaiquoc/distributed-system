package main

import (
	"testing"
)

func TestKeyValueStore(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		value    string
		expected string
	}{
		{
			name:     "Test SET operation",
			key:      "key1",
			value:    "value1",
			expected: "value1",
		},
		{
			name:     "Test GET operation for a non-existent key",
			key:      "key2",
			value:    "",
			expected: "",
		},
		{
			name:     "Test SET operation with an existing key",
			key:      "key1",
			value:    "value2",
			expected: "value2",
		},
	}

	// Create a new instance of the key-value store, initializing data
	kv := NewKVStore()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Test SET operation
			if test.value != "" {
				kv.Set(test.key, test.value)
			}

			// Test GET operation
			value, err := kv.Get(test.key)
			if err != nil {
				t.Errorf("Error getting value for %s: %v", test.key, err)
			}
			if value != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, value)
			}
		})
	}
}
