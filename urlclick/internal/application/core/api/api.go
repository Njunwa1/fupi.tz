package api

import (
	"context"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/clicks"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) GetUrlClicksAggregates(ctx context.Context, request *clicks.UserUrlRequest) (*clicks.UrlClicksAggregates, error) {
	_, err := a.db.GetUrlClickAggregates(ctx, request)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
