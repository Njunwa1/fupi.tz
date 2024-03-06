package ports

import (
	"context"
	"github.com/Njunwa1/shortener/internal/application/core/domain"
)

type APIPort interface {
	CreateShortUrl(ctx context.Context, url domain.Url) (domain.Url, error)
}
