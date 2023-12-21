package storage

type MemoryStorage struct {
	
}

func NewMemoryStorage() *MemoryStorage {
	return &
}
func (s *MemoryStorage) Get(id int) *types.User {
	return &types.User {
		ID: 1,
		Name: "Foo",
	}
}