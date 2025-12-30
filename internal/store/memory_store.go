package store

import (
	"time"

	"github.com/sokchamrernheng/oncey/internal/errors"
)

type MemoryStore struct {
	data map[string]CachedData
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string]CachedData),
	}
}
func NewMemoryStores() MemoryStore {
	return MemoryStore{
		data: make(map[string]CachedData),
	}
}

func (ms *MemoryStore) Get(key string) ([]byte, error) {
	data, ok := ms.data[key]
	if !ok {
		return nil, errors.ErrNotFound
	}
	return data.Value, nil
}

func (ms *MemoryStore) Set(key string, value []byte, ttl time.Duration) error {
	ms.data[key] = CachedData{
		Status:    "sucess",
		Value:     value,
		ExpiresAt: time.Now().Add(ttl).Unix(),
	}
	return nil
}

func (ms *MemoryStore) Forget(key string) error {
	delete(ms.data, key)
	return nil
}
