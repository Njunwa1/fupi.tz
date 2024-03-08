package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/redirect/internal/application/core/domain"
)

type DBPort interface {
	SaveUrlClick(ctx context.Context, click domain.UrlClick) error
	GetUrlClicksByShortUrl(ctx context.Context, shortUrl string) ([]domain.UrlClick, error)
}
