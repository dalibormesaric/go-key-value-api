package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dalibormesaric/go-key-value-api/internal/history"
	"github.com/dalibormesaric/go-key-value-api/internal/store"
)

type httpServer struct {
	store   store.Store
	history history.History
}

type keyValue struct {
	Key   string
	Value string
}

func NewHTTPServer(port int) *http.Server {
	s := &httpServer{
		store:   *store.New(),
		history: *history.New(),
	}

	h := http.NewServeMux()
	h.HandleFunc("GET /store", s.getStore)
	h.HandleFunc("POST /store", s.postStore)
	h.HandleFunc("DELETE /store", s.deleteStore)
	h.HandleFunc("GET /history", s.getHistory)

	return &http.Server{
		Handler: h,
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
	}
}

func (s *httpServer) getStore(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	v := s.store.Get(key)
	if v == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%s", v)
}

func (s *httpServer) postStore(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var keyValue keyValue
	err = json.Unmarshal(b, &keyValue)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	s.store.Add(keyValue.Key, keyValue.Value)
	s.history.Append(keyValue.Key, keyValue.Value)
}

func (s *httpServer) deleteStore(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	s.store.Delete(key)
	s.history.Delete(key)
}

func (s *httpServer) getHistory(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	v := s.history.Get(key)
	if v == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	result, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%v", string(result))
}
