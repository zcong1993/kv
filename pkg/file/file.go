package file

import (
	"fmt"

	"go.etcd.io/bbolt"
)

type Store struct {
	name       string
	bucketName string
	filePath   string
	db         *bbolt.DB
}

func NewFileStore(name string, bucketName string, filePath string) *Store {
	db, err := bbolt.Open(filePath, 0666, nil)

	if err != nil {
		panic(err)
	}

	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucket([]byte(bucketName))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return &Store{
		name:       name,
		bucketName: bucketName,
		filePath:   filePath,
		db:         db,
	}
}

func (s *Store) Name() string {
	return s.name
}

func (s *Store) Put(key string, value []byte) error {
	err := s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(s.bucketName))
		return b.Put([]byte(key), value)
	})
	return err
}

func (s *Store) Get(key string) (value []byte, err error) {
	err = s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(s.bucketName))
		v := b.Get([]byte(key))
		value = v
		return nil
	})
	return
}

func (s *Store) Delete(key string) error {
	err := s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(s.bucketName))
		return b.Delete([]byte(key))
	})
	return err
}
