package ports

import (
	"context"
	"github.com/Njunwa1/keygen/internal/application/core/domain"
)

type APIPort interface {
	GenerateShortUrlKey(ctx context.Context) (domain.KeyGenLogEntry, error)
}
