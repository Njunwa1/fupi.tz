package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/domain"
	"github.com/Njunwa1/fupitz-proto/golang/url"
)

type APIPort interface {
	CreateShortUrl(ctx context.Context, url domain.Url) (domain.Url, error)
	GetUrlByShortUrl(ctx context.Context, shortUrl string) (*url.UrlResponse, error)
	GetAllUserUrls(ctx context.Context, request *url.UserUrlsRequest) (*url.UserUrlsResponse, error)
}
