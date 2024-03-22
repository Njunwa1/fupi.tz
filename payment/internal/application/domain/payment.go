package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Payment represents a payment transaction related to a subscription.
type Payment struct {
	ID             primitive.ObjectID // Unique identifier for the payment
	SubscriptionID uint               // ID of the subscription
	Amount         float64            // Payment amount
	Currency       string             // Currency of the payment
	PaymentMethod  string             // Payment Method used
	PaymentDate    time.Time          // Date and time of the payment
}
