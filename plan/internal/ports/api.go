package ports

import "context"

type APIPort interface {
	CreatePlan(ctx context.Context)
	GetPlans(ctx context.Context)
}
