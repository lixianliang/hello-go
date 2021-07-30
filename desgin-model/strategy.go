package main

import (
	"fmt"
	"os"
)

type StorageStrategy interface {
	Save(name string, data []byte) error
}

var strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := strategys[t]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}

	return s, nil
}

type fileStorage struct{}

func (s *fileStorage) Save(name string, data []byte) error {
	return os.WriteFile(name, data, os.ModeAppend)
}

type encryptFileStorage struct{}

func (s *encryptFileStorage) Save(name string, data []byte) error {
	data, err := encrypt(data)
	if err != nil {
		return err
	}

	return os.WriteFile(name, data, os.ModeAppend)
}

func encrypt(data []byte) ([]byte, error) {
	return data, nil
}

func main() {
	data, sensitive := getData()
	strategyType := "file"
	if sensitive {
		strategyType = "encrypt_file"
	}

	storage, _ := NewStorageStrategy(strategyType)
	storage.Save("./test.txt", data)
}

func getData() ([]byte, bool) {
	return []byte("test data"), false
}
