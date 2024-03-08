package api

import (
	"context"
	"github.com/Njunwa1/fupi.tz/redirect/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/redirect/internal/ports"
)

type Application struct {
	db        ports.DBPort
	shortener ports.ShortenerPort
}

func NewApplication(db ports.DBPort, shortener ports.ShortenerPort) *Application {
	return &Application{db: db, shortener: shortener}
}

func (a *Application) CreateUrlClick(ctx context.Context, shortUrl string, click domain.UrlClick) (domain.UrlClick, error) {
	res, err := a.shortener.GetUrlByShortKey(ctx, shortUrl)
	if err != nil {
		return domain.UrlClick{}, err
	}
	click.UrlID = res.Id
	err = a.db.SaveUrlClick(ctx, click)
	if err != nil {
		return domain.UrlClick{}, err
	}
	if err != nil {
		return domain.UrlClick{}, err
	}
	return click, nil
}
