package ports

import (
	"context"
	"github.com/njunwa1/fupi.tz/shortener/internal/application/core/domain"
)

type APIPort interface {
	CreateShortUrl(ctx context.Context, url domain.Url) (domain.Url, error)
}
