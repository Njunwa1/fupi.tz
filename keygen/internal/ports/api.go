package ports

import (
	"context"
	"fupi.tz/keygen/internal/application/core/domain"
)

type APIPort interface {
	GenerateShortUrlKey(ctx context.Context) (domain.KeyGenLogEntry, error)
}
