package grpc

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/subscription"
)

func (a Adapter) CreateSubscription(ctx context.Context, request *subscription.SubscriptionRequest) (*subscription.SubscriptionResponse, error) {
	result, err := a.api.CreateSubscription(ctx, request)
	if err != nil {
		return &subscription.SubscriptionResponse{}, err
	}
	return result, nil
}

func
