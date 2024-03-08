package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/redirect/internal/application/core/domain"
)

type APIPort interface {
	CreateUrlClick(ctx context.Context, shortUrl string, click domain.UrlClick) (domain.UrlClick, error)
}
