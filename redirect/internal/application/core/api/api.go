package api

import (
	"context"
	"errors"
	"github.com/Njunwa1/fupi.tz/redirect/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/clicks"
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/metadata"
	"log/slog"
)

type Application struct {
	shortener ports.ShortenerPort
	redis     ports.RedisPort
	rabbitmq  ports.RabbitMQPort
}

func NewApplication(
	shortener ports.ShortenerPort,
	redis ports.RedisPort,
	rabbitmq ports.RabbitMQPort,
) *Application {
	return &Application{shortener: shortener, redis: redis, rabbitmq: rabbitmq}
}

func (a *Application) CreateUrlClick(ctx context.Context, request *clicks.UrlClickRequest, md metadata.MD) (*clicks.UrlClickResponse, error) {

	// 1. get the url by short key from redis cache
	var dbRes *url.UrlResponse
	var returnRes *clicks.UrlClickResponse
	res, err := a.redis.GetUrl(ctx, request.GetShortUrl())
	if err != nil {
		if errors.Is(err, redis.Nil) {
			// 2. if not in cache get it from db then set cache
			slog.Info("Url not available in reddis cache, fetching from db")
			dbRes, err := a.shortener.GetUrlByShortKey(ctx, request.ShortUrl)
			if err != nil {
				return &clicks.UrlClickResponse{}, err
			}
			_ = a.redis.SetUrl(ctx, request.ShortUrl, dbRes) // save the url in redis
		} else {
			return &clicks.UrlClickResponse{}, err
		}
	} else {
		slog.Info("Url fetched from redis cache", "response", res)
	}
	if dbRes != nil {
		request.UrlID = dbRes.Id
		returnRes = &clicks.UrlClickResponse{
			Id:          dbRes.Id,
			WebUrl:      dbRes.WebUrl,
			ShortUrl:    dbRes.Short,
			AndroidUrl:  dbRes.AndroidUrl,
			IosUrl:      dbRes.IosUrl,
			ExpiryAt:    dbRes.ExpiryAt,
			CustomAlias: dbRes.CustomAlias,
			Password:    dbRes.Password,
			Type:        dbRes.Type,
		}
	}
	if res != nil {
		request.UrlID = res.Id
		returnRes = &clicks.UrlClickResponse{
			Id:          res.Id,
			WebUrl:      res.WebUrl,
			ShortUrl:    res.ShortUrl,
			AndroidUrl:  res.AndroidUrl,
			IosUrl:      res.IosUrl,
			ExpiryAt:    res.ExpiryAt,
			CustomAlias: res.CustomAlias,
			Password:    res.Password,
			Type:        res.Type,
		}
	}

	slog.Info("Publishing message to RabbitMQ", "request", request, "metadata", md)
	// 3. Send details to the message broker for saving the click
	err = a.rabbitmq.PublishClickEvent(ctx, request, md)
	if err != nil {
		slog.Error("Failed to publish click event", "err", err)
	}

	return returnRes, nil
}
