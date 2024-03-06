package ports

import (
	"context"
	"github.com/njunwa1/fupi.tz/keygen/internal/application/core/domain"
)

type DBPort interface {
	SaveShortUrlKey(ctx context.Context, entry domain.KeyGenLogEntry) error
}
