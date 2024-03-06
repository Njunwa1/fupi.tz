package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/keygen/internal/application/core/domain"
)

type DBPort interface {
	SaveShortUrlKey(ctx context.Context, entry domain.KeyGenLogEntry) error
}
