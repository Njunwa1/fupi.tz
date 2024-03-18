package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Njunwa1/fupitz-proto/golang/redirect"
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"github.com/redis/go-redis/v9"
	"log"
	"log/slog"
	"time"
)

type Adapter struct {
	Client *redis.Client
}

func NewAdapter(dataSourceUrl string) *Adapter {
	client := redis.NewClient(&redis.Options{
		Addr:     dataSourceUrl,
		Password: "",
		DB:       0,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("failed to connect to redis: %s", err)
		return nil
	}
	return &Adapter{Client: client}
}

func (a *Adapter) GetUrl(ctx context.Context, shortKey string) (*redirect.RedirectResponse, error) {
	val, err := a.Client.HGet(ctx, shortKey, shortKey).Result()

	if err != nil {
		return nil, err
	}

	var response redirect.RedirectResponse
	err = json.Unmarshal([]byte(val), &response)
	if err != nil {
		log.Fatalf("Error while unmarshaling %s", err)
	}

	return &response, nil
}

func (a *Adapter) SetUrl(ctx context.Context, shortKey string, url *url.UrlResponse) error {

	//Unlike GetUrl, SetUrl receives data from fn that returns url.CreateUrlResponse
	urlData := &redirect.RedirectResponse{
		Id: url.Id, ShortUrl: url.Short, WebUrl: url.WebUrl, IosUrl: url.IosUrl,
		AndroidUrl: url.AndroidUrl, Password: url.Password, ExpiryAt: url.ExpiryAt,
		CustomAlias: url.CustomAlias, Type: url.Type,
	}
	jsonData, err := json.Marshal(urlData)
	if err != nil {
		slog.Error("Error occured while marshalling data", "err", err)
	}
	err = a.Client.HSet(ctx, shortKey, shortKey, jsonData).Err()
	if err != nil {
		return err
	}
	// Set expiration for the entire hash
	expiration := 10 * time.Minute
	err = a.Client.Expire(ctx, shortKey, expiration).Err()
	if err != nil {
		slog.Error("Error while setting expiry day", "err", err)
	}
	return nil
}
