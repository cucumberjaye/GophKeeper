package main

import (
	"github.com/cucumberjaye/GophKeeper/internal/pkg/app/clientapp"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.FatalLevel)
	clientApp, err := clientapp.New()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	err = clientApp.Run()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
