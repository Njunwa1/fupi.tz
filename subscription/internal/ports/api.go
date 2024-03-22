package ports

import "context"

type APIPort interface {
	CreateSubscription(ctx context.Context)
	CreatePayment(ctx context.Context)
	GetPlans(ctx context.Context)
	GetSubscriptions(ctx context.Context)
	GetUserSubscriptionPlan(ctx context.Context)
}
