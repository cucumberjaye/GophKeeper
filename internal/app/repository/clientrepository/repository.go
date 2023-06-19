package clientrepository

import "github.com/redis/go-redis/v9"

// ClientStorage - структура для хранения данных клинта.
type ClientStorage struct {
	rdb *redis.Client
}

// New - инициализирует структуру ClientStorage.
func New(rdb *redis.Client) *ClientStorage {
	return &ClientStorage{
		rdb: rdb,
	}
}
