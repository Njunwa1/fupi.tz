package domain

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

// Errors
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// Payload represents the payload of the Paseto token
type Payload struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Role       string    `json:"role"`
	IssuedAt   time.Time `json:"iat"`
	Expiration time.Time `json:"exp"`
}

// NewPayload creates a new payload for the Paseto token
func NewPayload(username, role string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Payload{
		ID:         tokenID.String(),
		Username:   username,
		Role:       role,
		IssuedAt:   time.Now(),
		Expiration: time.Now().Add(duration),
	}, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.Expiration) {
		return ErrExpiredToken
	}
	return nil
}
