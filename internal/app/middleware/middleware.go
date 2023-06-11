package middleware

import (
	"context"

	"github.com/cucumberjaye/GophKeeper/internal/pkg/tokens"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthenticationGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	notAuth := []string{
		"/pb.Authentication/Registration",
		"/pb.Authentication/Authentication",
	}

	for _, val := range notAuth {
		if val == info.FullMethod {
			return handler(ctx, req)
		}
	}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		values := md.Get("authentication")
		if len(values) > 0 {
			authToken := values[0]

			id, err := tokens.ParseToken(authToken)
			if err == nil {
				ctx = metadata.AppendToOutgoingContext(ctx, "user_id", id)
				return handler(ctx, req)
			}
			log.Error().Err(err).Send()
		}
	}
	return nil, status.Error(codes.Unauthenticated, "unauthenticated")

}
