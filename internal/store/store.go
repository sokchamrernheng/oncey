package store

// Store defines the persistence layer for idempotency keys.
//
// Implementations MUST ensure:
// - Get returns (nil, false) if the key does not exist or is expired
// - Expired keys MUST NOT be returned
// - Set stores the key for at least the provided TTL
// - Forget removes the key immediately
type Store interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl int) error
	Forget(key string) error
}

type CachedData struct {
	Id string
	Status string
	Value []byte
	TTL string
}