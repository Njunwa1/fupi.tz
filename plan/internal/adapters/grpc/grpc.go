package grpc

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/plan"
)

func (a Adapter) GetPlan(ctx context.Context, request *plan.PlanByIdRequest) (*plan.PlanResponse, error) {
	result, err := a.api.GetPlan(ctx, request)
	if err != nil {
		return &plan.PlanResponse{}, err
	}
	return result, nil
}

func (a Adapter) GetPlans(ctx context.Context) ([]*plan.PlanResponse, error) {
	result, err := a.api.GetPlans(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}
