package api

import (
	"context"
	"fmt"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/shortener/internal/ports"
	"log"
)

type Application struct {
	db     ports.DBPort
	keygen ports.KeyGenPort
}

func NewApplication(db ports.DBPort, keygen ports.KeyGenPort) *Application {
	return &Application{db: db, keygen: keygen}
}

func (a *Application) CreateShortUrl(ctx context.Context, url domain.Url) (domain.Url, error) {
	log.Println("Sending Original URL to shortening service", url)
	url.Short, _ = a.keygen.GenerateShortUrlKey(ctx)
	log.Println("Generated short url", url.Short)
	err := a.db.SaveUrl(ctx, url)
	if err != nil {
		log.Println("Failed to save url to database: ", err)
		return domain.Url{}, err
	}
	return url, nil
}

func (a *Application) GetUrlByShortUrl(ctx context.Context, shortUrl string) (domain.Url, error) {
	url, err := a.db.GetUrlByShortUrl(ctx, shortUrl)
	fmt.Println("URL from database", url)
	if err != nil {
		log.Println("Failed to get url from database: ", err)
		return domain.Url{}, err
	}
	return url, nil
}
