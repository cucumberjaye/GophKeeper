package main

import "github.com/cucumberjaye/GophKeeper/internal/pkg/app/serverapp"

func main() {
	serverApp, err := serverapp.New()
	if err != nil {
		panic(err)
	}

	err = serverApp.Run()
	if err != nil {
		panic(err)
	}
}
