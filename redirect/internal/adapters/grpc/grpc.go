package grpc

import (
	"context"
	"github.com/Njunwa1/fupi.tz/redirect/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/redirect/proto/golang/clicks"
)

func (a Adapter) CreateUrlClick(ctx context.Context, request *clicks.UrlClickRequest) (*clicks.UrlClickResponse, error) {
	click := domain.NewUrlClick(
		request.UrlID,
		request.UserAgent,
		request.IpAddress,
		request.Referrer,
		request.DeviceType,
		request.Os,
		request.Browser,
		request.Country,
		request.City,
		float64(request.Latitude),
		float64(request.Longitude),
	)
	click, err := a.api.CreateUrlClick(ctx, request.ShortUrl, click)
	if err != nil {
		return nil, err
	}
	return &clicks.UrlClickResponse{}, nil
}
