package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr string

	DBConnConfig *DBConnConfig
}

type DBConnConfig struct {
	User         string
	Password     string
	Host         string
	Port         string
	DBName       string
	PoolMaxConns string
}

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
	}, nil
}

func getEnvIfExists(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
