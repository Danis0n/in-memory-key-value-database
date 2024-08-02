package engine

import "sync"

var HashTableBuilder = func() hashTable { return NewHashTable() }

type HashTable struct {
	mutex sync.RWMutex
	table map[string]string
}

func NewHashTable() *HashTable {
	return &HashTable{
		table: make(map[string]string),
	}
}

func (h *HashTable) Set(key, value string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	h.table[key] = value
}

func (h *HashTable) Get(key string) (string, bool) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	data, found := h.table[key]

	return data, found
}

func (h *HashTable) Del(key string) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	delete(h.table, key)
}
