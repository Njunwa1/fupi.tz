package ports

import (
	"context"
	"github.com/njunwa1/fupi.tz/shortener/internal/application/core/domain"
)

type DBPort interface {
	SaveUrl(context.Context, domain.Url) error
}
