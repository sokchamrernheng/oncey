package store

import "time"

// Store defines the persistence layer for idempotency keys.
//
// Implementations MUST ensure:
// - Get returns (nil, false) if the key does not exist or is expired
// - Expired keys MUST NOT be returned
// - Set stores the key for at least the provided TTL
// - Forget removes the key immediately
type Store interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl time.Duration) error
	Forget(key string) error
}

type CachedData struct {
	Status    string
	Value     []byte
	ExpiresAt int64
}
