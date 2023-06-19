package serverhandler

import (
	"context"

	"github.com/cucumberjaye/GophKeeper/internal/app/models"
	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
)

// AuthServer - струтура для использования grpc хендлеров.
type AuthServer struct {
	pb.UnimplementedAuthenticationServer

	Service KeeperService // Service - содержит методы слоя с логикой.
}

// Registration - grpc метод для регистрации.
func (s *AuthServer) Registration(ctx context.Context, in *pb.RegistrationRequest) (*pb.ResponseStatus, error) {
	err := validator.New().Struct(&models.LoginPasswordValidate{
		Login:    in.Login,
		Password: in.Password,
	})
	if err != nil {
		return &pb.ResponseStatus{Status: pb.ResponseStatus_FAIL}, err
	}

	if err := s.Service.AddUser(in.Login, in.Password); err != nil {
		log.Error().Err(err).Send()
		return &pb.ResponseStatus{Status: pb.ResponseStatus_FAIL}, err
	}

	return &pb.ResponseStatus{Status: pb.ResponseStatus_OK}, nil
}

// Authentication - grpc метод для аутентификации.
func (s *AuthServer) Authentication(ctx context.Context, in *pb.AuthenticationRequest) (*pb.AuthToken, error) {
	token, err := s.Service.CreateToken(in.Login, in.Password)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return &pb.AuthToken{Token: token}, nil
}
