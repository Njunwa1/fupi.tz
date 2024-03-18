package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/redirect"
	"github.com/Njunwa1/fupitz-proto/golang/url"
)

type RedisPort interface {
	GetUrl(ctx context.Context, shortKey string) (*redirect.RedirectResponse, error)
	SetUrl(ctx context.Context, shortKey string, url *url.UrlResponse) error
}
