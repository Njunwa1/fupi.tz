package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/subscription"
)

type APIPort interface {
	CreateSubscription(ctx context.Context, request *subscription.SubscriptionRequest) (*subscription.SubscriptionResponse, error)
	GetUserActiveSubscriptions(ctx context.Context, request *subscription.UserActiveSubscriptionRequest) (*subscription.SubscriptionResponse, error)
}
