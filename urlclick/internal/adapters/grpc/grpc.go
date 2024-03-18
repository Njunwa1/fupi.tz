package grpc

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/clicks"
)

func (a Adapter) GetUserUrlWithClicks(ctx context.Context, request *clicks.UserUrlRequest) (*clicks.UrlClicksAggregates, error) {
	_, err := a.api.GetUrlClicksAggregates(ctx, request)
	if err != nil {
		return nil, err
	}
	return &clicks.UrlClicksAggregates{}, nil
}
