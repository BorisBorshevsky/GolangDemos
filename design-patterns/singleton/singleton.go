package singleton

import (
	"fmt"
	"sync"
)

var (
	r *repo
)

type repo struct {
	items map[string]string
	mu    sync.RWMutex
}

func (r *repo) Set(key, data string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[key] = data
}

func (r *repo) Get(key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.items[key]
	if !ok {
		return "", fmt.Errorf("The '%s' is not presented", key)
	}
	return item, nil
}

func Repo() *repo {
	if r == nil {
		r = &repo{
			items: make(map[string]string),
		}
	}
	return r
}
