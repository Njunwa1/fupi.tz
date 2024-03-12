package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UrlClick struct {
	UrlID      *primitive.ObjectID `bson:"urlid,omitempty"`
	UserAgent  string
	IPAddress  string
	Referrer   string  // Referring URL
	DeviceType string  // Type of device (e.g., desktop, mobile)
	Browser    string  // Web browser used for the click
	OS         string  // Operating system of the user
	Country    string  // Country of the user
	City       string  // City of the user
	Latitude   float64 // Latitude of the user's location
	Longitude  float64 // Longitude of the user's location
	CreatedAt  time.Time
}

func NewUrlClick(urlID primitive.ObjectID, userAgent, ipAddress, referrer, deviceType, browser, os, country, city string, latitude, longitude float64) UrlClick {
	return UrlClick{
		UrlID:      &urlID,
		UserAgent:  userAgent,
		IPAddress:  ipAddress,
		Referrer:   referrer,
		DeviceType: deviceType,
		Browser:    browser,
		OS:         os,
		Country:    country,
		City:       city,
		Latitude:   latitude,
		Longitude:  longitude,
		CreatedAt:  time.Now(),
	}
}
