package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config - хранит все переменные конфигурации.
type Config struct {
	ServerAddr string

	DBConnConfig *DBConnConfig
	RedisConfig  *RedisConfig
}

// DBConnConfig - хранит переменные конфигурации для PostgreSQL.
type DBConnConfig struct {
	User         string
	Password     string
	Host         string
	Port         string
	DBName       string
	PoolMaxConns string
}

// RedisConfig - хранит переменные конфигурации для Redis.
type RedisConfig struct {
	Host     string
	Port     string
	Password string
}

// New - загружает переменные окружения из dev.env в процесс и считывает переменные окружения в структуру Config.
func New() (*Config, error) {
	if err := godotenv.Load("dev.env"); err != nil {
		return nil, fmt.Errorf("load env file failed with error: %w", err)
	}

	return &Config{
		ServerAddr: getEnvIfExists("SERVER_ADDR", ":3000"),
		DBConnConfig: &DBConnConfig{
			User:         getEnvIfExists("DB_USER", "postgres"),
			Password:     getEnvIfExists("DB_PASSWORD", "qwerty1234"),
			Host:         getEnvIfExists("DB_HOST", "postgres"),
			Port:         getEnvIfExists("DB_PORT", "5432"),
			DBName:       getEnvIfExists("DB_NAME", "postgres"),
			PoolMaxConns: getEnvIfExists("POOL_MAX_CONNS", "5"),
		},
		RedisConfig: &RedisConfig{
			Host:     getEnvIfExists("REDIS_HOST", "127.0.0.1"),
			Port:     getEnvIfExists("REDIS_PORT", "6379"),
			Password: getEnvIfExists("REDIS_PASSWORD", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81"),
		},
	}, nil
}

// getEnvIfExists - считывает переменную окружения по key, если такой нет, возвращает значение по умолчанию.
func getEnvIfExists(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
