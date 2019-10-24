package kv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPut(t *testing.T, store Store) {
	ast := assert.New(t)
	key := "test"
	val := []byte("val")

	v, err := store.Get(key)

	ast.Nil(err)
	ast.Nil(v)

	err = store.Put(key, val)
	ast.Nil(err)

	v, err = store.Get(key)
	ast.Nil(err)
	ast.Equal(val, v)
}

func TestGet(t *testing.T, store Store) {
	ast := assert.New(t)
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

func TestDelete(t *testing.T, store Store) {
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
