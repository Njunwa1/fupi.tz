package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/application/core/domain"
	"github.com/Njunwa1/fupitz-proto/golang/clicks"
)

type DBPort interface {
	SaveUrlClick(ctx context.Context, click domain.UrlClick) error
	GetUrlClickAggregates(ctx context.Context, request *clicks.UserUrlRequest) (*clicks.UrlClicksAggregates, error)
}
