package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/plan"
)

type APIPort interface {
	CreatePlan(ctx context.Context) (*plan.PlanResponse, error)
	GetPlans(ctx context.Context) ([]*plan.PlanResponse, error)
	GetPlan(ctx context.Context, request *plan.PlanByIdRequest) (*plan.PlanResponse, error)
}
