package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/url"
)

type ShortenerPort interface {
	GetUrlByShortKey(ctx context.Context, shortKey string) (*url.UrlResponse, error)
}
