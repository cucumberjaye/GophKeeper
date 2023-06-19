package redis

import (
	"github.com/cucumberjaye/GophKeeper/configs"
	"github.com/redis/go-redis/v9"
)

// New - создает подключение к Redis.
func New(cfg *configs.RedisConfig) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
	})

	return rdb
}
