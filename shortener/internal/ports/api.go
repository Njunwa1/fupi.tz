package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/url"
)

type APIPort interface {
	CreateShortUrl(ctx context.Context, request *url.UrlRequest) (*url.UrlResponse, error)
	GetUrlByShortUrl(ctx context.Context, shortUrl string) (*url.UrlResponse, error)
	GetAllUserUrls(ctx context.Context, request *url.UserUrlsRequest) (*url.UserUrlsResponse, error)
}
