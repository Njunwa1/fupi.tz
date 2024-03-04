package api

import (
	"context"
	"fupi.tz/shortener/internal/application/core/domain"
	"fupi.tz/shortener/internal/ports"
)

type Application struct {
	db     ports.DBPort
	keygen ports.KeyGenPort
}

func NewApplication(db ports.DBPort, keygen ports.KeyGenPort) *Application {
	return &Application{db: db, keygen: keygen}
}

func (a *Application) CreateShortUrl(ctx context.Context, url domain.Url) (domain.Url, error) {
	url.Short, _ = a.keygen.GenerateShortUrlKey(ctx)
	err := a.db.SaveUrl(ctx, url)
	if err != nil {
		return domain.Url{}, err
	}
	return url, nil
}
