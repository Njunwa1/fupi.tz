package grpc

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"log/slog"
)

func (a Adapter) CreateShortUrl(ctx context.Context, request *url.UrlRequest) (*url.UrlResponse, error) {
	slog.Info("Create Url request", "request", request)
	result, err := a.api.CreateShortUrl(ctx, request)
	if err != nil {
		return nil, err
	}
	return &url.UrlResponse{
		Id:    result.Id,
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
