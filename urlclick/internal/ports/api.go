package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/clicks"
)

type APIPort interface {
	GetUrlClicksAggregates(ctx context.Context, request *clicks.UserUrlRequest) (*clicks.UrlClicksAggregates, error)
}
