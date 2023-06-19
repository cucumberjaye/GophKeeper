package serverapp

import (
	"fmt"
	"net"

	"github.com/cucumberjaye/GophKeeper/configs"
	"github.com/cucumberjaye/GophKeeper/internal/app/handler/serverhandler"
	"github.com/cucumberjaye/GophKeeper/internal/app/middleware"
	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"github.com/cucumberjaye/GophKeeper/internal/app/repository/serverrepository"
	"github.com/cucumberjaye/GophKeeper/internal/app/service"
	"github.com/cucumberjaye/GophKeeper/pkg/postgres.go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// App - структура для запуска сервера.
type App struct {
	gs *grpc.Server
}

// New - инициализирует структуру App.
func New() (*App, error) {
	cfg, err := configs.New()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	pool, err := postgres.New(cfg.DBConnConfig)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	repo := serverrepository.New(pool)

	svc := service.New(repo)

	gs := grpc.NewServer(grpc.UnaryInterceptor(middleware.AuthenticationGRPC))
	pb.RegisterAuthenticationServer(gs, &serverhandler.AuthServer{
		Service: svc,
	})
	pb.RegisterStorageServer(gs, &serverhandler.StorageServer{
		Service: svc,
	})
	reflection.Register(gs)

	return &App{gs: gs}, nil
}

// Run - запускает сервер.
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
