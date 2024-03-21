package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

type KVRestService struct {
	kvStore *KVStore
}

func NewKVRestService() *KVRestService {
	return &KVRestService{kvStore: NewKVStore()}
}

func (kvRestService *KVRestService) handleGet(w http.ResponseWriter, r *http.Request) {

	key := chi.URLParam(r, "key")
	// logging the key
	fmt.Println("received key: ", key)
	// logging full url
	fmt.Println("received url: ", r.URL.String())
	value, err := kvRestService.kvStore.Get(key)
	if err != nil {
		if errors.Is(err, errKeyNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			panic(err)
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}

func (kvRestService *KVRestService) handlePut(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	value := string(body)
	kvRestService.kvStore.Set(key, value)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(value))
}

func (kvRestService *KVRestService) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Get("/api/kv/{key}", kvRestService.handleGet)
	mux.Put("/api/kv/{key}", kvRestService.handlePut)
	return mux
}

// Create a new router
// Define the GET and PUT endpoints
// Start the server
func main() {
	kvRestService := NewKVRestService()
	http.ListenAndServe(":8080", kvRestService.routes())
}
