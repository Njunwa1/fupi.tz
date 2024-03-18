package ports

import (
	"github.com/Njunwa1/fupi.tz/auth/internal/application/core/domain"
	"time"
)

type PasetoPort interface {
	CreateToken(userID, role string, duration time.Duration) (string, error)
	VerifyToken(token string) (*domain.Payload, error)
}
