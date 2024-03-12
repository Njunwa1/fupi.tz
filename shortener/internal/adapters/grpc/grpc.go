package grpc

import (
	"context"
	"github.com/Njunwa1/fupi.tz-proto/golang/url"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
	"time"
)

func (a Adapter) Create(ctx context.Context, request *url.CreateUrlRequest) (*url.CreateUrlResponse, error) {
	slog.Info("Create Url request", "request", request)
	urlType := domain.UrlType{Name: request.Type}
	expiryAt, _ := time.Parse("2006-01-02", request.ExpiryAt)
	newUrl := domain.NewUrl(
		urlType,
		request.CustomAlias,
		request.Password,
		request.QrcodeUrl,
		request.WebUrl,
		request.AppleUrl,
		request.AndroidUrl,
		primitive.ObjectID{},
		expiryAt,
	) //returns an address
	result, err := a.api.CreateShortUrl(ctx, *newUrl)
	if err != nil {
		return nil, err
	}
	return &url.CreateUrlResponse{
		Id:    result.Id.Hex(),
		Short: result.Short,
	}, nil
}

func (a Adapter) GetUrlByKey(ctx context.Context, request *url.GetUrlByKeyRequest) (*url.CreateUrlResponse, error) {
	slog.Info("Get Url request", "request", request)
	result, err := a.api.GetUrlByShortUrl(ctx, request.ShortUrl)
	if err != nil {
		return nil, err
	}
	return &url.CreateUrlResponse{
		Id:          result.Id.Hex(),
		Type:        result.UrlType.Name,
		WebUrl:      result.WebUrl,
		AppleUrl:    result.IOSUrl,
		AndroidUrl:  result.AndroidUrl,
		Short:       result.Short,
		ExpiryAt:    result.ExpiryAt.Format("2006-01-02"),
		CustomAlias: result.CustomAlias,
		Password:    result.Password,
	}, nil
}
