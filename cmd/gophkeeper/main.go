package main

import "github.com/cucumberjaye/GophKeeper/internal/pkg/app/serverapp"

func main() {
	a, err := serverapp.New()
	if err != nil {
		panic(err)
	}

	err = a.Run()
	if err != nil {
		panic(err)
	}
}
