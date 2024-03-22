package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Features struct {
	Links          int8
	QrCodes        int8
	LinkInBio      int8
	CustomAlias    int8
	CustomDomains  int8
	LinkPassword   bool
	LinkExpiration bool
	LinkStats      bool
	LinkStatsStore time.Duration
	UTMBuilder     bool
	CustomQrCodes  bool
}

// Plan represents a subscription plan offered by the SaaS platform.
type Plan struct {
	ID          primitive.ObjectID //Unique identifier for the plan
	Name        string             //Name of the plan
	Description string             // Description of the plan
	Price       float64            // Price of the plan
	Features    Features           // Features of the plan
}
