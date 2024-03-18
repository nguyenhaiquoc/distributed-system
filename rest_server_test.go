package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Write Test for handleGet
// Test that the handleGet function returns the correct value
// for a given key
func TestHandleGet(t *testing.T) {
	fmt.Println("This goes to standard output.")
	kvRestService := NewKVRestService()
	kvRestService.kvStore.Set("test", "value")
	req := httptest.NewRequest("GET", "/api/kv/test", nil)
	w := httptest.NewRecorder()

	kvRestService.routes().ServeHTTP(w, req)

	resp := w.Result()
	// print response

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
	if string(body) != "value" {
		t.Errorf("Expected value %s, got %s", "value", string(body))
	}
}
