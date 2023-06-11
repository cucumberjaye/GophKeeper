package serverhandler

import (
	"context"

	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"github.com/rs/zerolog/log"
)

type AuthServer struct {
	pb.UnimplementedAuthenticationServer

	Service KeeperService
}

func (s *AuthServer) Registration(ctx context.Context, in *pb.RegistrationRequest) (*pb.ResponseStatus, error) {
	if err := s.Service.AddUser(in.Login, in.Password); err != nil {
		log.Error().Err(err).Send()
		return &pb.ResponseStatus{Status: pb.ResponseStatus_FAIL}, err
	}

	return &pb.ResponseStatus{Status: pb.ResponseStatus_OK}, nil
}

func (s *AuthServer) Authentication(ctx context.Context, in *pb.AuthenticationRequest) (*pb.AuthToken, error) {
	token, err := s.Service.CreateToken(in.Login, in.Password)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}

	return &pb.AuthToken{Token: token}, nil
}
