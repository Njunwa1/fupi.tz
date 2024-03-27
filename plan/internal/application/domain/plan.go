package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Features represents the features offered in a subscription plan.
type Features struct {
	Links          int8          `json:"links" bson:"links"`                       // Number of links
	QrCodes        int8          `json:"qr_codes" bson:"qr_codes"`                 // Number of QR codes
	LinkInBio      int8          `json:"link_in_bio" bson:"link_in_bio"`           // Link in bio feature
	CustomAlias    int8          `json:"custom_alias" bson:"custom_alias"`         // Custom alias feature
	CustomDomains  int8          `json:"custom_domains" bson:"custom_domains"`     // Custom domains feature
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

func NewPlan(name, description string, price float32, features Features) *Plan {
	return &Plan{
		Name:        name,
		Description: description,
		Price:       price,
		Features:    features,
	}
}

var Plans = []Plan{
	{
		Name:        "Free",
		Description: "Ideal for individuals and small businesses getting started.",
		Price:       0.0,
		Features: Features{
			Links:          3,
			QrCodes:        3,
			LinkInBio:      1,
			CustomAlias:    0,
			CustomDomains:  0,
			LinkPassword:   false,
			LinkExpiration: false,
			LinkStats:      true,
			UTMBuilder:     false,
			CustomQrCodes:  false,
			LinkStatsStore: time.Duration(30) * 24 * time.Hour,
		},
	},
	{
		Name:        "Standard",
		Description: "Great for growing businesses looking for more customization options.",
		Price:       9.99,
		Features: Features{
			Links:          10,
			QrCodes:        3,
			LinkInBio:      3,
			CustomAlias:    3,
			CustomDomains:  1,
			LinkPassword:   true,
			LinkExpiration: true,
			LinkStats:      true,
			UTMBuilder:     false,
			CustomQrCodes:  false,
			LinkStatsStore: time.Duration(365) * 24 * time.Hour,
		},
	},
	{
		Name:        "Premium",
		Description: "For advanced users and businesses requiring comprehensive analytics and branding.",
		Price:       19.99,
		Features: Features{
			Links:          100,
			QrCodes:        100,
			LinkInBio:      100,
			CustomAlias:    100,
			CustomDomains:  100,
			LinkPassword:   true,
			LinkExpiration: true,
			LinkStats:      true,
			UTMBuilder:     true,
			CustomQrCodes:  true,
			LinkStatsStore: time.Duration(365) * 24 * time.Hour,
		},
	},
}
