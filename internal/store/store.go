package store

import "sync"

type Store struct {
	mu       sync.RWMutex
	keyvalue map[string]string
}

func New() *Store {
	return &Store{
		keyvalue: make(map[string]string),
	}
}

func (s *Store) Add(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.keyvalue[key] = value
}

func (s *Store) Get(key string) string {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.keyvalue[key]
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.keyvalue, key)
}
