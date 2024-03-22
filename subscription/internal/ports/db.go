package ports

import "context"

type DBPort interface {
	SaveSubscription(ctx context.Context)
	SavePayment(ctx context.Context)
	GetPlans(ctx context.Context)
	GetSubscriptions(ctx context.Context)
	GetUserSubscriptionPlan(ctx context.Context)
}
