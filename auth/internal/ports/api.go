package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/auth/internal/application/core/domain"
	"github.com/Njunwa1/fupitz-proto/golang/user"
)

type APIPort interface {
	Login(ctx context.Context, email string, password string) (string, error)
	Register(ctx context.Context, request *user.RegisterRequest) (*user.RegisterResponse, error)
	VerifyToken(token string) (*domain.Payload, error)
}
