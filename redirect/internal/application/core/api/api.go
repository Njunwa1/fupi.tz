package api

import (
	"context"
	"github.com/Njunwa1/fupi.tz-proto/golang/clicks"
	"github.com/Njunwa1/fupi.tz-proto/golang/url"
	"github.com/Njunwa1/fupi.tz/redirect/internal/ports"
	"google.golang.org/grpc/metadata"
	"log"
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

func (a *Application) CreateUrlClick(ctx context.Context, request *clicks.UrlClickRequest, md metadata.MD) (*url.CreateUrlResponse, error) {
	// 1. get the url by short key from redis cache
	redisRes, err := a.redis.GetUrl(ctx, request.ShortUrl)
	if err != nil {
		log.Printf("Error while getting url from redis: %s", err)
	} else {
		return redisRes, nil
	}

	// 2. if not in cache get it from db then set cache
	res, err := a.shortener.GetUrlByShortKey(ctx, request.ShortUrl)
	if err != nil {
		return &url.CreateUrlResponse{}, err
	}
	_ = a.redis.SetUrl(ctx, request.ShortUrl, res) // save the url in redis
	request.UrlID = res.Id                         // set the url id

	// 3. Send details to the message broker for saving the click
	err = a.rabbitmq.PublishClickEvent(ctx, request, md)
	if err != nil {
		slog.Error("Failed to publish click event", "err", err)
	}
	return res, nil
}
