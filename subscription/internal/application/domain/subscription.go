package domain

import (
	"github.com/Njunwa1/fupitz-proto/golang/plan"
	"strconv"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Features represents the features offered in a subscription plan.
type Features struct {
	Links          int32         `json:"links" bson:"links"`                       // Number of links
	QrCodes        int32         `json:"qr_codes" bson:"qr_codes"`                 // Number of QR codes
	LinkInBio      int32         `json:"link_in_bio" bson:"link_in_bio"`           // Link in bio feature
	CustomAlias    int32         `json:"custom_alias" bson:"custom_alias"`         // Custom alias feature
	CustomDomains  int32         `json:"custom_domains" bson:"custom_domains"`     // Custom domains feature
	LinkPassword   bool          `json:"link_password" bson:"link_password"`       // Link password feature
	LinkExpiration bool          `json:"link_expiration" bson:"link_expiration"`   // Link expiration feature
	LinkStats      bool          `json:"link_stats" bson:"link_stats"`             // Link stats feature
	LinkStatsStore time.Duration `json:"link_stats_store" bson:"link_stats_store"` // Duration to store link stats
	UTMBuilder     bool          `json:"utm_builder" bson:"utm_builder"`           // UTM builder feature
	CustomQrCodes  bool          `json:"custom_qr_codes" bson:"custom_qr_codes"`   // Custom QR codes feature
}

// Plan represents a subscription plan offered by the SaaS platform.
type Plan struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`        // Unique identifier for the plan
	Name        string             `json:"name" bson:"name"`               // Name of the plan
	Description string             `json:"description" bson:"description"` // Description of the plan
	Price       float32            `json:"price" bson:"price"`             // Price of the plan
	Features    Features           `json:"features" bson:"features"`       // Features of the plan
}

// Subscription represents a user's subscription to a plan.
type Subscription struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`      // Unique identifier for the subscription
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`       // ID of the user
	Plan      Plan               `json:"plan" bson:"plan"`             // ID of the plan
	StartDate time.Time          `json:"start_date" bson:"start_date"` // Start date of the subscription
	EndDate   time.Time          `json:"end_date" bson:"end_date"`     // End date of the subscription
	IsActive  bool               `json:"is_active" bson:"is_active"`   // Indicates if the subscription is active
}

func NewSubscription(userId primitive.ObjectID, plan Plan, duration time.Duration) *Subscription {
	return &Subscription{
		UserID:    userId,
		Plan:      plan,
		StartDate: time.Now(),
		EndDate:   time.Now().Add(duration * 24 * time.Hour),
		IsActive:  false,
	}
}

func PlanFromResponse(res *plan.PlanResponse) Plan {
	planId, _ := primitive.ObjectIDFromHex(res.GetId())
	price, _ := strconv.Atoi(res.Price)
	return Plan{
		ID:          planId,
		Name:        res.Name,
		Description: res.Description,
		Price:       float32(price),
		Features: Features{
			//Links:          res.Features.links,
			QrCodes:        res.Features.QrCodes,
			LinkInBio:      res.Features.LinkInBio,
			CustomAlias:    res.Features.CustomAlias,
			CustomDomains:  res.Features.CustomDomains,
			LinkPassword:   res.Features.LinkPassword,
			LinkExpiration: res.Features.LinkExpiration,
			LinkStats:      res.Features.LinkStats,
			LinkStatsStore: time.Duration(res.Features.LinkStatsStore),
			UTMBuilder:     res.Features.UtmBuilder,
			CustomQrCodes:  res.Features.CustomQrCodes,
		},
	}
}

func (s Subscription) TotalPrice() float32 {
	duration := s.EndDate.Sub(s.StartDate)
	return s.Plan.Price * float32(duration*24*time.Hour)
}
