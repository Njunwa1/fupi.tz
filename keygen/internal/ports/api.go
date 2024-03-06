package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/keygen/internal/application/core/domain"
)

type APIPort interface {
	GenerateShortUrlKey(ctx context.Context) (domain.KeyGenLogEntry, error)
}
