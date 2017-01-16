package main

import (
	"sync"
)

type cMapLockInt struct {
	store map[int]int
	mu    *sync.RWMutex
}

func (cl *cMapLockInt) set(key, val int) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	cl.store[key] = val
}

func (cl *cMapLockInt) get(key int) int {
	cl.mu.RLock()
	defer cl.mu.RUnlock()

	return cl.store[key]
}

func (cl *cMapLockInt) del(key int) {
	cl.mu.Lock()
	defer cl.mu.Unlock()

	delete(cl.store, key)
}

func newCMapLockInt() *cMapLockInt {
	return &cMapLockInt{
		store: make(map[int]int),
		mu:    &sync.RWMutex{},
	}
}
