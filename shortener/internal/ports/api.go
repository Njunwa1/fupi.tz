package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/domain"
)

type APIPort interface {
	CreateShortUrl(ctx context.Context, url domain.Url) (domain.Url, error)
	GetUrlByShortUrl(ctx context.Context, shortUrl string) (domain.Url, error)
}
