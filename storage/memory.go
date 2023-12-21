package storage

import "github.com/EAS-Kalem/moto_go/types"

type MemoryStorage struct {
	
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}
func (s *MemoryStorage) Get(id int) *types.User {
	return &types.User {
		ID: 1,
		Name: "Foo",
	}
}