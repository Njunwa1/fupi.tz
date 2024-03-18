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
		slog.Error("Create Url request failed", "error", err)
		return nil, err
	}
	return result, nil
}

func (a Adapter) GetUrlByKey(ctx context.Context, request *url.UrlByKeyRequest) (*url.UrlResponse, error) {
	slog.Info("Get Url request", "request", request)
	result, err := a.api.GetUrlByShortUrl(ctx, request.ShortUrl)
	if err != nil {
		slog.Error("Get Url request failed", "error", err)
		return nil, err
	}
	return result, nil
}

func (a Adapter) GetAllUserUrls(ctx context.Context, request *url.UserUrlsRequest) (*url.UserUrlsResponse, error) {
	urlsResponse, err := a.api.GetAllUserUrls(ctx, request)
	if err != nil {
		slog.Error("Get all user urls request failed", "error", err)
		return nil, err
	}
	return urlsResponse, nil
}
