package domain

import (
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Subscription represents a user's subscription to a plan.
type Subscription struct {
	ID        primitive.ObjectID // Unique identifier for the subscription
	UserID    primitive.ObjectID // ID of the user
	PlanID    primitive.ObjectID // ID of the plan
	StartDate time.Time          // Start date of the subscription
	EndDate   time.Time          // End date of the subscription
	IsActive  bool               // Indicates if the subscription is active
}
