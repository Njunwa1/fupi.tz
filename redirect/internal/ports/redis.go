package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/clicks"
	"github.com/Njunwa1/fupitz-proto/golang/url"
)

type RedisPort interface {
	GetUrl(ctx context.Context, shortKey string) (*clicks.UrlClickResponse, error)
	SetUrl(ctx context.Context, shortKey string, url *url.UrlResponse) error
}
