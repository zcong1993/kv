package memory_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zcong1993/kv/pkg/memory"
)

func TestMemStore_Get(t *testing.T) {
	ast := assert.New(t)
	store := memory.NewMemStore("test")

	key := "test"
	val := []byte("val")

	_ = store.Put(key, val)
	v, err := store.Get(key)
	ast.Nil(err)
	ast.Equal(val, v)

	v, err = store.Get("nil")
	ast.Nil(err)
	ast.Nil(v)
}

func TestMemStore_Put(t *testing.T) {
	store := memory.NewMemStore("test")
	key := "test"
	val := []byte("val")

	v, err := store.Get(key)

	ast := assert.New(t)
	ast.Nil(err)
	ast.Nil(v)

	err = store.Put(key, val)
	ast.Nil(err)

	v, err = store.Get(key)
	ast.Nil(err)
	ast.Equal(val, v)
}

func TestMemStore_Delete(t *testing.T) {
	store := memory.NewMemStore("test")
	ast := assert.New(t)
	key := "test"
	val := []byte("val")
	_ = store.Put(key, val)

	err := store.Delete(key)
	ast.Nil(err)

	v, err := store.Get(key)
	ast.Nil(err)
	ast.Nil(v)
}
