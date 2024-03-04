package ports

import (
	"context"
	"fupi.tz/shortener/internal/application/core/domain"
)

type DBPort interface {
	SaveUrl(context.Context, domain.Url) error
}
