package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/subscription/internal/application/domain"
)

type DBPort interface {
	CreateSubscription(ctx context.Context, subscription domain.Subscription) (domain.Subscription, error)
	GetUserActiveSubscription(ctx context.Context, useId string) (domain.Subscription, error)
}
