package database

import "flareAPI/types"

type MemoryStorage struct {
}

func NewLocalDB() (*MemoryStorage, error) {
	return &MemoryStorage{}, nil
}

func (m *MemoryStorage) GetAdmin(id int) *types.User {
	return &types.User{ID: id, Name: "LocalAdmin"}
}

func (m *MemoryStorage) Close() {
	return
}
