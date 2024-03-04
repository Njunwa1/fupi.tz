package ports

import (
	"context"
	"fupi.tz/keygen/internal/application/core/domain"
)

type DBPort interface {
	SaveShortUrlKey(ctx context.Context, entry domain.KeyGenLogEntry) error
}
