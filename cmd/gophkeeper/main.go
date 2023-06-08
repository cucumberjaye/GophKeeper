package main

import "github.com/cucumberjaye/GophKeeper/pkg/grpc"

func main() {
	a, err := grpc.New()
	if err != nil {
		panic(err)
	}

	err = a.Run()
	if err != nil {
		panic(err)
	}
}
