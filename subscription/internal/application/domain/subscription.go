package domain

import (
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Subscription represents a user's subscription to a plan.
type Subscription struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`      // Unique identifier for the subscription
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`       // ID of the user
	PlanID    primitive.ObjectID `json:"plan_id" bson:"plan_id"`       // ID of the plan
	StartDate time.Time          `json:"start_date" bson:"start_date"` // Start date of the subscription
	EndDate   time.Time          `json:"end_date" bson:"end_date"`     // End date of the subscription
	IsActive  bool               `json:"is_active" bson:"is_active"`   // Indicates if the subscription is active
}

func NewSubscription(userId, planId primitive.ObjectID, duration time.Duration) *Subscription {
	return &Subscription{
		UserID:    userId,
		PlanID:    planId,
		StartDate: time.Now(),
		EndDate:   time.Now().Add(duration * 24 * time.Hour),
		IsActive:  false,
	}
}
