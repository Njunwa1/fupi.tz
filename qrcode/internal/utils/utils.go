package utils

import (
	"errors"
	"github.com/Njunwa1/fupi.tz/qrcode/config"
	"github.com/o1egl/paseto"
	"time"
)

// Errors
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

// UserIDKey is the key for the user id in the context
type UserIDKey struct{}

// Payload represents the payload of the Paseto token
type Payload struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Role       string    `json:"role"`
	IssuedAt   time.Time `json:"iat"`
	Expiration time.Time `json:"exp"`
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.Expiration) {
		return ErrExpiredToken
	}
	return nil
}

func VerifyToken(token string) (*Payload, error) {
	p := paseto.NewV2()
	payload := &Payload{}
	err := p.Decrypt(token, config.GetSymmetricKey(), payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	if err := payload.Valid(); err != nil {
		return nil, ErrExpiredToken
	}
	return payload, nil
}

func Contains(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}
