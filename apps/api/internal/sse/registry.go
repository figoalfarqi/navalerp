package sse

import "sync"

type Registry struct {
	mu   sync.Mutex
	hubs map[string]*Hub
}

func NewRegistry() *Registry {
	return &Registry{
		hubs: make(map[string]*Hub),
	}
}

func (r *Registry) GetHub(role, page string) *Hub {
	key := role + ":" + page

	r.mu.Lock()
	defer r.mu.Unlock()

	if hub, ok := r.hubs[key]; ok {
		return hub
	}

	hub := NewHub()
	r.hubs[key] = hub
	return hub
}
