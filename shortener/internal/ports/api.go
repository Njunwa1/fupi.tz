package ports

import (
	"context"
	"fupi.tz/shortener/internal/application/core/domain"
)

type APIPort interface {
	CreateShortUrl(ctx context.Context, url domain.Url) (domain.Url, error)
}
