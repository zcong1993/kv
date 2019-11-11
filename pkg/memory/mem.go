package memory

type MemStore struct {
	name  string
	store map[string][]byte
}

func NewMemStore(name string) *MemStore {
	return &MemStore{
		name:  name,
		store: make(map[string][]byte, 0),
	}
}

func (ms *MemStore) Name() string {
	return ms.name
}

func (ms *MemStore) Get(key string) ([]byte, error) {
	val, ok := ms.store[key]
	if !ok {
		return nil, nil
	}
	return val, nil
}

func (ms *MemStore) Put(key string, value []byte) error {
	ms.store[key] = value
	return nil
}

func (ms *MemStore) Delete(key string) error {
	delete(ms.store, key)
	return nil
}
