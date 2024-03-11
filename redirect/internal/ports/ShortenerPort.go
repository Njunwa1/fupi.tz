package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz-proto/golang/url"
)

type ShortenerPort interface {
	GetUrlByShortKey(ctx context.Context, shortKey string) (*url.CreateUrlResponse, error)
}
