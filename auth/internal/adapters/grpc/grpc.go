package grpc

import (
	"context"
	"errors"
	"github.com/Njunwa1/fupitz-proto/golang/user"
	"google.golang.org/grpc/metadata"
)

func (a Adapter) Login(ctx context.Context, request *user.LoginRequest) (*user.LoginResponse, error) {
	// md, _ := metadata.FromIncomingContext(ctx)
	res, err := a.api.Login(ctx, request)
	if err != nil {
		// slog.Error("Failed to create url click", "err", err)
		return &user.LoginResponse{}, err
	}
	return res, nil
}

func (a Adapter) Register(ctx context.Context, request *user.RegisterRequest) (*user.RegisterResponse, error) {
	res, err := a.api.Register(ctx, request)
	if err != nil {
		return &user.RegisterResponse{}, err
	}
	return res, nil
}

func (a Adapter) VerifyToken(ctx context.Context, request *user.VerifyTokenRequest) (*user.VerifyTokenResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	token := md.Get("authorization")
	if len(token) == 0 {
		return &user.VerifyTokenResponse{}, errors.New("token not found")
	}
	_, err := a.api.VerifyToken(token[0])
	if err != nil {
		return &user.VerifyTokenResponse{}, err
	}
	return user.VerifyTokenResponse{accessToken: token}, nil
}
