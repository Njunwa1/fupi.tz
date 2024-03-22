package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

//import "go.mongodb.org/mongo-driver/bson/primitive"

// Plan represents a subscription plan offered by the SaaS platform.
type Plan struct {
	ID          primitive.ObjectID //Unique identifier for the plan
	Name        string             //Name of the plan
	Description string             // Description of the plan
	Price       float64            // Price of the plan
}
