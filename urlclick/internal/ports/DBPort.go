package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/application/core/domain"
)

type DBPort interface {
	SaveUrlClick(ctx context.Context, click domain.UrlClick) error
}
