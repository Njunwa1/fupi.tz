package api

import (
	"context"
	"fupi.tz/keygen/internal/application/core/domain"
	"fupi.tz/keygen/internal/application/core/utils"
	"fupi.tz/keygen/internal/ports"
	"time"
)

const keyOutputLength = 7

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

// GenerateShortUrlKey implements the APIPort interface
func (a *Application) GenerateShortUrlKey(ctx context.Context) (domain.KeyGenLogEntry, error) {
	entry := domain.KeyGenLogEntry{
		ShortUrl:  utils.GenerateKey(keyOutputLength),
		CreatedAt: time.Now(),
	}
	err := a.db.SaveShortUrlKey(ctx, entry)
	if err != nil {
		return domain.KeyGenLogEntry{}, err
	}
	return entry, nil
}
