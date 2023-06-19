package clientapp

import (
	"github.com/cucumberjaye/GophKeeper/configs"
	"github.com/cucumberjaye/GophKeeper/internal/app/handler/clienthandler"
	"github.com/cucumberjaye/GophKeeper/internal/app/repository/clientrepository"
	"github.com/cucumberjaye/GophKeeper/pkg/redis"
	"github.com/rs/zerolog/log"
)

// ClientApp - структура для запуска клиента.
type ClientApp struct {
	client *clienthandler.KeeperClient
}

// New - инициализирует структуру ClientApp.
func New() (*ClientApp, error) {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	rdb := redis.New(cfg.RedisConfig)
	repo := clientrepository.New(rdb)

	clinet, err := clienthandler.New(repo)

	return &ClientApp{client: clinet}, nil
}

// Run - запускает клиента.
func (a *ClientApp) Run() error {
	log.Info().Msg("Client starting...")

	return a.client.Run()
}
