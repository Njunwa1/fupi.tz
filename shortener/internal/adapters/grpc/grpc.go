package grpc

import (
	"context"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/domain"
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
	"time"
)

func (a Adapter) CreateShortUrl(ctx context.Context, request *url.UrlRequest) (*url.UrlResponse, error) {
	slog.Info("Create Url request", "request", request)
	urlType := domain.UrlType{Name: request.Type}
	expiryAt, _ := time.Parse("2006-01-02", request.ExpiryAt)
	newUrl := domain.NewUrl(
		urlType,
		request.CustomAlias,
		request.Password,
		request.QrcodeUrl,
		request.WebUrl,
		request.IosUrl,
		request.AndroidUrl,
		primitive.ObjectID{},
		expiryAt,
	) //returns an address
	result, err := a.api.CreateShortUrl(ctx, *newUrl)
	if err != nil {
		return nil, err
	}
	return &url.UrlResponse{
		Id:    result.Id.Hex(),
		Short: result.Short,
	}, nil
}

func (a Adapter) GetUrlByKey(ctx context.Context, request *url.UrlByKeyRequest) (*url.UrlResponse, error) {
	slog.Info("Get Url request", "request", request)
	result, err := a.api.GetUrlByShortUrl(ctx, request.ShortUrl)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (a Adapter) GetAllUserUrls(ctx context.Context, request *url.UserUrlsRequest) (*url.UserUrlsResponse, error) {
	urlsResponse, err := a.api.GetAllUserUrls(ctx, request)
	if err != nil {
		return nil, err
	}
	return urlsResponse, nil
}
