package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/domain"
)

type DBPort interface {
	SaveUrl(context.Context, domain.Url) error
}
