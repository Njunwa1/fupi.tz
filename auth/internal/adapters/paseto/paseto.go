package paseto

import (
	"fmt"
	"github.com/Njunwa1/fupi.tz/auth/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/auth/internal/ports"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (ports.PasetoPort, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

func (p *PasetoMaker) CreateToken(userID, role string, duration time.Duration) (string, error) {
	payload, err := domain.NewPayload(userID, role, duration)
	if err != nil {
		return "", err
	}
	return p.paseto.Encrypt(p.symmetricKey, payload, nil)
}

func (p *PasetoMaker) VerifyToken(token string) (*domain.Payload, error) {
	payload := &domain.Payload{}
	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)
	if err != nil {
		return nil, domain.ErrInvalidToken
	}
	if err := payload.Valid(); err != nil {
		return nil, domain.ErrExpiredToken
	}
	return payload, nil
}
