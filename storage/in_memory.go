package storage

import "sync"

type InMemoryStorage struct {
	cache map[string]bool
	mutex *sync.RWMutex
}

func newInMemoryStorage() InMemoryStorage {
	cache := make(map[string]bool)
	return InMemoryStorage{
		cache: cache,
		mutex: &sync.RWMutex{},
	}
}

func (m InMemoryStorage) Get(key string) bool {
	return m.cache[key]
}

func (m InMemoryStorage) Put(key string, value bool) {
	m.mutex.Lock()
	m.cache[key] = value
	m.mutex.Unlock()
}

func (m InMemoryStorage) Delete(key string) {
	delete(m.cache, key)
}

func (m InMemoryStorage) GetAndDelKey() string {
	key := ""

	if len(m.cache) == 0 {
		return key
	}

	m.mutex.RLock()
	for k, _ := range m.cache {
		key = k
		break
	}
	m.Delete(key)
	m.mutex.RUnlock()

	return key
}
