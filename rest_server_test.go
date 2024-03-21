package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Write Test for handleGet
// Test that the handleGet function returns the correct value
// for a given key
func TestHandleGet(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		expected   string
		statusCode int
	}{
		{name: "Existing Key", key: "test-key", expected: "exits-value", statusCode: http.StatusOK},
		{name: "Non-Existing Key", key: "nonexistent", expected: "", statusCode: http.StatusNotFound},
		{name: "Existing Key Empty Value", key: "empty", expected: "", statusCode: http.StatusOK},
	}

	kvRestService := NewKVRestService()
	kvRestService.kvStore.Set("test-key", "exits-value")
	kvRestService.kvStore.Set("empty", "")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/kv/"+tt.key, nil)
			w := httptest.NewRecorder()

			kvRestService.routes().ServeHTTP(w, req)
			resp := w.Result()
			body, _ := io.ReadAll(resp.Body)
			if resp.StatusCode != tt.statusCode {
				t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
			}
			if resp.StatusCode == http.StatusOK && string(body) != tt.expected {
				t.Errorf("Expected value = %s, got = %s", tt.expected, string(body))
			}
		})
	}
}

func TestHandlePut(t *testing.T) {
	tests := []struct {
		name       string
		key        string
		value      string
		statusCode int
	}{
		{name: "New Key", key: "new-key", value: "new-value", statusCode: http.StatusOK},
		{name: "Update Key", key: "new-key", value: "updated-value", statusCode: http.StatusOK},
	}

	kvRestService := NewKVRestService()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("PUT", "/api/kv/"+tt.key, nil)
			req.Body = io.NopCloser(strings.NewReader(tt.value))
			w := httptest.NewRecorder()
			kvRestService.routes().ServeHTTP(w, req)
			resp := w.Result()
			if resp.StatusCode != tt.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.statusCode, resp.StatusCode)
			}

			// Check if the value was set correctly
			value, err := kvRestService.kvStore.Get(tt.key)
			if err != nil {
				t.Error("Error getting value from store")
			}
			if value != tt.value {
				t.Errorf("Expected value = %s, got = %s", tt.value, value)
			}

		})
	}
}
