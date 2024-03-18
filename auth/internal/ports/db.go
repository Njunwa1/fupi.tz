package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/auth/internal/application/core/domain"
)

type DBPort interface {
	SaveUser(ctx context.Context, user domain.User) error
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
}
