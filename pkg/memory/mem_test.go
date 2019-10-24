package memory_test

import (
	"testing"

	"github.com/zcong1993/kv/pkg/kv"

	"github.com/zcong1993/kv/pkg/memory"
)

func TestMemStore_Get(t *testing.T) {
	store := memory.NewMemStore("test")
	kv.TestGet(t, store)
}

func TestMemStore_Put(t *testing.T) {
	store := memory.NewMemStore("test")
	kv.TestPut(t, store)
}

func TestMemStore_Delete(t *testing.T) {
	store := memory.NewMemStore("test")
	kv.TestDelete(t, store)
}
