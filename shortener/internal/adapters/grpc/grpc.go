package grpc

import (
	"context"
	"github.com/njunwa1/fupi.tz/proto/golang/url"
	"github.com/njunwa1/fupi.tz/shortener/internal/application/core/domain"
	"log/slog"
	"time"
)

func (a Adapter) Create(ctx context.Context, request *url.CreateUrlRequest) (*url.CreateUrlResponse, error) {
	slog.Info("Create Url request", "request", request)
	urlType := domain.UrlType{Name: request.Type}
	expiryAt, _ := time.Parse("2006-01-02", request.ExpiryAt)
	newUrl := domain.NewUrl(
		1,
		urlType,
		request.CustomAlias,
		request.Password,
		request.QrcodeUrl,
		request.Original,
		expiryAt,
	) //returns an address
	result, err := a.api.CreateShortUrl(ctx, *newUrl)
	if err != nil {
		return nil, err
	}
	return &url.CreateUrlResponse{Short: result.Short}, nil
}
