package grpc

import (
	"context"
	"errors"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log/slog"
)

var accessibleRoles = map[string][]string{
	"/UrlClicks/GetUserUrlWithClicks": {"user"},
}

func AuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	slog.Info("AuthInterceptor", "info", info.FullMethod)
	roles, ok := accessibleRoles[info.FullMethod]
	if !ok {
		return handler(ctx, req)
	}
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
	if !utils.Contains(roles, payload.Role) {
		return nil, status.Errorf(codes.PermissionDenied, "you don't have permission to access this resource")
	}
	return handler(ctx, req)
}
