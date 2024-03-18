package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/auth/internal/application/core/domain"
)

type APIPort interface {
	Login(ctx context.Context, email string, password string) (string, error)
	Register(ctx context.Context, user domain.User) (domain.User, error)
	VerifyToken(token string) (*domain.Payload, error)
}
