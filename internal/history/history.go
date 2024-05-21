package history

import (
	"sync"
	"time"
)

type History struct {
	mu   sync.Mutex
	data map[string][]historyKeyValue
}

type historyKeyValue struct {
	Value     string
	Timestamp time.Time
	State     string
}

const (
	added   = "added"
	updated = "updated"
	deleted = "deleted"
)

func New() *History {
	return &History{
		data: make(map[string][]historyKeyValue),
	}
}

func (h *History) Get(key string) []historyKeyValue {
	return h.data[key]
}

func (h *History) Append(key, value string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	v, ok := h.data[key]
	state := updated
	if !ok || v[0].Value == "" {
		state = added
	}
	h.data[key] = append([]historyKeyValue{{
		Value:     value,
		Timestamp: time.Now(),
		State:     state,
	}}, h.data[key]...)
}

func (h *History) Delete(key string) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.data[key] = append([]historyKeyValue{{
		Value:     "",
		Timestamp: time.Now(),
		State:     deleted,
	}}, h.data[key]...)
}
