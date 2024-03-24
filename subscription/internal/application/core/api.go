package core

import (
	"context"
	"github.com/Njunwa1/fupi.tz/subscription/internal/application/domain"
	"github.com/Njunwa1/fupi.tz/subscription/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/subscription"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) CreateSubscription(ctx context.Context, request *subscription.SubscriptionRequest) (*subscription.SubscriptionResponse, error) {
	duration := 30 * 24 * time.Hour // 30 days
	userId, _ := primitive.ObjectIDFromHex(request.GetUserId())
	planId, _ := primitive.ObjectIDFromHex(request.GetPlanId())
	newSub := domain.NewSubscription(
		userId, planId, duration,
	)
	sub, err := a.db.CreateSubscription(ctx, *newSub)
	if err != nil {
		return &subscription.SubscriptionResponse{}, err
	}
	return &subscription.SubscriptionResponse{
		Id:     sub.ID.Hex(),
		UserId: sub.PlanID.Hex(),
		PlanId: sub.PlanID.Hex(),
	}, nil
}

func (a *Application) GetUserActiveSubscriptions(ctx context.Context, request *subscription.UserActiveSubscriptionRequest) (*subscription.SubscriptionResponse, error) {
	sub, err := a.db.GetUserActiveSubscription(ctx, request.GetUserId())
	if err != nil {
		return &subscription.SubscriptionResponse{}, err
	}
	return &subscription.SubscriptionResponse{
		Id:     sub.ID.Hex(),
		UserId: sub.PlanID.Hex(),
		PlanId: sub.PlanID.Hex(),
	}, nil
}
