package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/plan"
)

type PlanPort interface {
	GetPlan(ctx context.Context, planId string) (*plan.PlanResponse, error)
}
