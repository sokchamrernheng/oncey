package store

type MemoryStore struct {
	data map[string][]byte
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[string][]byte),
	};
}
func NewMemoryStores() MemoryStore {
	return MemoryStore{
		data: make(map[string][]byte),
	};
}


func (ms *MemoryStore) Get(key string) ([]byte, error) {
	data := []byte{1, 2, 3}
	return data,nil;
}

func (ms *MemoryStore) Set(key string, value []byte, ttl int) error {
	ms.data["hello"] = []byte("World")
	return nil;
}

func (ms *MemoryStore) Forget(key string) error {
	return nil;
}