package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DBPort interface {
	SaveUrl(context.Context, domain.Url) error
	GetUrlByShortUrl(context.Context, string) (domain.Url, error)
	GetAllUserUrls(context.Context, *primitive.ObjectID) ([]domain.Url, error)
}
