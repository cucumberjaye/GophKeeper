package grpc

import (
	"fmt"
	"net"

	"github.com/cucumberjaye/GophKeeper/configs"
	"github.com/cucumberjaye/GophKeeper/internal/handler"
	"github.com/cucumberjaye/GophKeeper/internal/middleware"
	"github.com/cucumberjaye/GophKeeper/internal/pb"
	"github.com/cucumberjaye/GophKeeper/internal/repository"
	"github.com/cucumberjaye/GophKeeper/internal/service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	gs *grpc.Server
}

func New() (*App, error) {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	repo, err := repository.New(cfg.DBConnConfig)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	svc := service.New(repo)

	gs := grpc.NewServer(grpc.UnaryInterceptor(middleware.AuthenticationGRPC))
	pb.RegisterAuthenticationServer(gs, &handler.AuthServer{
		Service: svc,
	})
	pb.RegisterStorageServer(gs, &handler.StorageServer{
		Service: svc,
	})
	reflection.Register(gs)

	return &App{gs: gs}, nil
}

func (a *App) Run() error {
	fmt.Println("server starting...")

	listen, err := net.Listen("tcp", ":3000")
	if err != nil {
		return fmt.Errorf("listen failed with error: %w", err)
	}

	if err := a.gs.Serve(listen); err != nil {
		return fmt.Errorf("grpc serve failed with error: %w", err)
	}

	return nil
}
