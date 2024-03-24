package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Payment represents a payment transaction related to a subscription.
type Payment struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`                // Unique identifier for the payment
	SubscriptionID primitive.ObjectID `json:"subscription_id" bson:"subscription_id"` // ID of the subscription
	Amount         float32            `json:"amount" bson:"amount"`                   // Payment amount
	Currency       string             `json:"currency" bson:"currency"`               // Currency of the payment
	PaymentMethod  string             `json:"payment_method" bson:"payment_method"`   // Payment Method used
	PaymentDate    time.Time          `json:"payment_date" bson:"payment_date"`       // Date and time of the payment
}

func NewPayment(subId primitive.ObjectID, currency, paymentMethod string, amount float32) *Payment {
	return &Payment{
		SubscriptionID: subId,
		Amount:         amount,
		Currency:       currency,
		PaymentMethod:  paymentMethod,
		PaymentDate:    time.Now(),
	}
}
