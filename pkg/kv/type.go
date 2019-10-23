package kv

type Store interface {
	Name() string
	Put(key string, value []byte) error
	Get(key string) (value []byte, err error)
	Delete(key string) error
}
