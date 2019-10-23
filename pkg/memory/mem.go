package memory

import "sync"

type MemStore struct {
	name  string
	store *sync.Map
}

func NewMemStore(name string) *MemStore {
	return &MemStore{
		name:  name,
		store: &sync.Map{},
	}
}

func (ms *MemStore) Name() string {
	return ms.name
}

func (ms *MemStore) Get(key string) ([]byte, error) {
	val, ok := ms.store.Load(key)
	if !ok {
		return nil, nil
	}
	return val.([]byte), nil
}

func (ms *MemStore) Put(key string, value []byte) error {
	ms.store.Store(key, value)
	return nil
}

func (ms *MemStore) Delete(key string) error {
	ms.store.Delete(key)
	return nil
}
