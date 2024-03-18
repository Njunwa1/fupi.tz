package grpc

import (
	"context"
	"errors"
	"github.com/Njunwa1/fupi.tz/shortener/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if len(md["authorization"]) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}
	payload, err := utils.VerifyToken(md["authorization"][0])
	if err != nil {
		if errors.Is(err, utils.ErrExpiredToken) {
			return nil, status.Errorf(codes.Unauthenticated, "token has expired")
		}
		return nil, status.Errorf(codes.Unauthenticated, "token is invalid")
	}
	ctx = context.WithValue(ctx, utils.UserIDKey{}, payload.UserID)
	return handler(ctx, req)
}
