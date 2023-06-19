package serverhandler

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/cucumberjaye/GophKeeper/internal/app/middleware"
	"github.com/cucumberjaye/GophKeeper/internal/app/pb"
	"github.com/cucumberjaye/GophKeeper/internal/app/service/mocks"
	"github.com/cucumberjaye/GophKeeper/internal/pkg/tokens"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

// HandlerSuite - струтрура содержащая поля для тестирования.
type HandlerSuite struct {
	suite.Suite

	authClient  pb.AuthenticationClient
	storeClient pb.StorageClient

	svc    *mocks.MockKeeperService
	closer func()

	authToken string
}

// TestHandlerSuite - создает HandlerSuite и запускает тесты.
func TestHandlerSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

// server - создает тестовый сервер и подключения к нему.
func (s *HandlerSuite) server() {
	ctx := context.Background()
	buffer := 1024 * 1024
	lis := bufconn.Listen(buffer)

	ctrl := gomock.NewController(s.T())

	m := mocks.NewMockKeeperService(ctrl)
	s.svc = m

	baseServer := grpc.NewServer(grpc.UnaryInterceptor(middleware.AuthenticationGRPC))
	pb.RegisterAuthenticationServer(baseServer, &AuthServer{Service: m})
	pb.RegisterStorageServer(baseServer, &StorageServer{Service: m})

	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Error().Err(err).Msg("error serving server")
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithInsecure())
	if err != nil {
		log.Error().Err(err).Msg("error connecting to server")
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Error().Err(err).Msg("error closing listener")
		}
		baseServer.Stop()
	}

	s.authClient = pb.NewAuthenticationClient(conn)
	s.storeClient = pb.NewStorageClient(conn)
	s.closer = closer
	authToken, err := tokens.CreateToken("test")
	s.Require().NoError(err)
	s.authToken = authToken
}

// TestRegistration - тестирует метод Registration.
func (s *HandlerSuite) TestRegistration() {
	s.server()
	defer s.closer()

	tests := map[string]struct {
		in  *pb.RegistrationRequest
		err error
	}{
		"OK": {
			in: &pb.RegistrationRequest{
				Login:    "11",
				Password: "22",
			},
			err: nil,
		},
		"Validation error": {
			in: &pb.RegistrationRequest{
				Login:    "",
				Password: "22",
			},
			err: errors.New("rpc error: code = Unknown desc = Key: 'LoginPasswordValidate.Login' Error:Field validation for 'Login' failed on the 'required' tag"),
		},
		"Service error": {
			in: &pb.RegistrationRequest{
				Login:    "11",
				Password: "33",
			},
			err: errors.New("rpc error: code = Unknown desc = service error"),
		},
	}

	for name, tt := range tests {
		s.Run(name, func() {
			if tt.err == nil {
				s.svc.EXPECT().AddUser(tt.in.Login, tt.in.Password).Return(nil)
			} else if name == "Service error" {
				s.svc.EXPECT().AddUser(tt.in.Login, tt.in.Password).Return(errors.New("service error"))
			}
			_, err := s.authClient.Registration(context.Background(), tt.in)
			if err != nil {
				s.Require().Equal(tt.err.Error(), err.Error())
			}
		})
	}
}

// TestAuthentication - тестирует метод Authentication.
func (s *HandlerSuite) TestAuthentication() {
	s.server()
	defer s.closer()

	tests := map[string]struct {
		in  *pb.AuthenticationRequest
		err error
	}{
		"OK": {
			in: &pb.AuthenticationRequest{
				Login:    "11",
				Password: "22",
			},
			err: nil,
		},
		"Service error": {
			in: &pb.AuthenticationRequest{
				Login:    "11",
				Password: "33",
			},
			err: errors.New("rpc error: code = Unknown desc = service error"),
		},
	}

	for name, tt := range tests {
		s.Run(name, func() {
			if tt.err == nil {
				s.svc.EXPECT().CreateToken(tt.in.Login, tt.in.Password).Return("", nil)
			} else if name == "Service error" {
				s.svc.EXPECT().CreateToken(tt.in.Login, tt.in.Password).Return("", errors.New("service error"))
			}
			_, err := s.authClient.Authentication(context.Background(), tt.in)
			if err != nil {
				s.Require().Equal(tt.err.Error(), err.Error())
			}
		})
	}
}
