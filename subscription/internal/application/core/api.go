package core

import (
	"context"
	"github.com/Njunwa1/fupi.tz/subscription/internal/application/domain"
	"github.com/Njunwa1/fupi.tz/subscription/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/subscription"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
	"time"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
	plan    ports.PlanPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, plan ports.PlanPort) *Application {
	return &Application{db: db, payment: payment, plan: plan}
}

func (a *Application) CreateSubscription(ctx context.Context, request *subscription.SubscriptionRequest) (*subscription.SubscriptionResponse, error) {
	duration := 30 * 24 * time.Hour // 30 days
	userId, _ := primitive.ObjectIDFromHex(request.GetUserId())
	planResponse, _ := a.plan.GetPlan(ctx, request.GetPlanId())
	plan := domain.PlanFromResponse(planResponse)
	newSub := domain.NewSubscription(
		userId, plan, duration,
	)
	sub, err := a.db.CreateSubscription(ctx, *newSub)
	if err != nil {
		return &subscription.SubscriptionResponse{}, err
	}
	_, err = a.payment.CreatePayment(ctx, sub)
	if err != nil {
		slog.Error("Could not make payment", "err", err)
		return nil, err
	}
	return &subscription.SubscriptionResponse{
		Id:     sub.ID.Hex(),
		UserId: sub.Plan.ID.Hex(),
		PlanId: sub.Plan.ID.Hex(),
	}, nil
}

func (a *Application) GetUserActiveSubscriptions(ctx context.Context, request *subscription.UserActiveSubscriptionRequest) (*subscription.SubscriptionResponse, error) {
	sub, err := a.db.GetUserActiveSubscription(ctx, request.GetUserId())
	if err != nil {
		return &subscription.SubscriptionResponse{}, err
	}
	return &subscription.SubscriptionResponse{
		Id:     sub.ID.Hex(),
		UserId: sub.Plan.ID.Hex(),
		PlanId: sub.Plan.ID.Hex(),
	}, nil
}
