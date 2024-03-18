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
	UserID     string    `json:"user_id"`
	Role       string    `json:"role"`
	IssuedAt   time.Time `json:"iat"`
	Expiration time.Time `json:"exp"`
}

// NewPayload creates a new payload for the Paseto token
func NewPayload(UserID, role string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	return &Payload{
		ID:         tokenID.String(),
		UserID:     UserID,
		Role:       role,
		IssuedAt:   time.Now(),
		Expiration: time.Now().Add(time.Hour * duration),
	}, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.Expiration) {
		return ErrExpiredToken
	}
	return nil
}
