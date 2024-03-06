package ports

import (
	"context"
	"github.com/njunwa1/fupi.tz/keygen/internal/application/core/domain"
)

type APIPort interface {
	GenerateShortUrlKey(ctx context.Context) (domain.KeyGenLogEntry, error)
}
