package core

import (
	"context"
	"github.com/Njunwa1/fupi.tz/plan/internal/application/domain"
	"github.com/Njunwa1/fupi.tz/plan/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/plan"
	"log/slog"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) CreatePlan(ctx context.Context) (*plan.PlanResponse, error) {
	err := a.db.Create(ctx, domain.Plans)
	if err != nil {
		slog.Info("Could not create plans")
		return &plan.PlanResponse{}, err
	}
	return &plan.PlanResponse{}, nil
}

func (a *Application) GetPlan(ctx context.Context, request *plan.PlanByIdRequest) (*plan.PlanResponse, error) {
	result, err := a.db.Get(ctx, request.GetId())
	if err != nil {
		return &plan.PlanResponse{}, err
	}
	return &plan.PlanResponse{
		Id:          result.ID.Hex(),
		Name:        result.Name,
		Description: result.Description,
		//Features: result.Features,
	}, nil
}

func (a *Application) GetPlans(ctx context.Context) ([]*plan.PlanResponse, error) {
	plans, err := a.db.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	var results []*plan.PlanResponse
	for _, p := range plans {
		planResponse := plan.PlanResponse{
			Id:          p.ID.Hex(),
			Name:        p.Name,
			Description: p.Description,
			//Price: p.Price,
			//Features: p.Features,
		}
		results = append(results, &planResponse)
	}
	return results, nil
}
