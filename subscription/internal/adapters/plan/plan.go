package plan

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/plan"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	plan plan.PlanClient
}

func NewAdapter(planServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(planServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := plan.NewPlanClient(conn)
	return &Adapter{plan: client}, nil
}

func (a *Adapter) GetPlan(ctx context.Context, planId string) (*plan.PlanResponse, error) {
	res, err := a.plan.GetPlan(ctx, &plan.PlanByIdRequest{
		Id: planId,
	})
	if err != nil {
		return &plan.PlanResponse{}, err
	}
	return res, nil
}
