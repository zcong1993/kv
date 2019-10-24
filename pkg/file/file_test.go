package file_test

import (
	"os"
	"testing"

	"github.com/zcong1993/kv/pkg/file"
	"github.com/zcong1993/kv/pkg/kv"
)

func unlink(filePath string) {
	_ = os.Remove(filePath)
}

func TestStore_Put(t *testing.T) {
	filePath := "./test_put.db"
	store := file.NewFileStore("test", filePath)
	kv.TestPut(t, store)
	unlink(filePath)
}

func TestStore_Get(t *testing.T) {
	filePath := "./test_get.db"
	store := file.NewFileStore("test", filePath)
	kv.TestGet(t, store)
	unlink(filePath)
}

func TestStore_Delete(t *testing.T) {
	filePath := "./test_delete.db"
	store := file.NewFileStore("test", filePath)
	kv.TestDelete(t, store)
	unlink(filePath)
}
