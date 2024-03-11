package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz-proto/golang/url"
)

type RedisPort interface {
	GetUrl(ctx context.Context, shortKey string) (*url.CreateUrlResponse, error)
	SetUrl(ctx context.Context, shortKey string, url *url.CreateUrlResponse) error
}
