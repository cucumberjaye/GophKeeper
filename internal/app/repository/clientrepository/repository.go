package clientrepository

import "github.com/redis/go-redis/v9"

type ClientStorage struct {
	rdb *redis.Client
}

func New(rdb *redis.Client) *ClientStorage {
	return &ClientStorage{
		rdb: rdb,
	}
}
