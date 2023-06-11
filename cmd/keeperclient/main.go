package main

import (
	"github.com/cucumberjaye/GophKeeper/internal/pkg/app/clientapp"
	"github.com/rs/zerolog/log"
)

func main() {
	a, err := clientapp.New()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	err = a.Run()
	if err != nil {
		log.Fatal().Err(err).Send()
	}
}
