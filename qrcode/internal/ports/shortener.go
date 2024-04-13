package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/url"
)

type ShortenerPort interface {
	CreateShortUrl(ctx context.Context, request *url.UrlRequest) (*url.UrlResponse, error)
}
