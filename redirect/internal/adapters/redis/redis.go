package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Njunwa1/fupi.tz-proto/golang/url"
	"github.com/redis/go-redis/v9"
	"log"
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

func (a *Adapter) GetUrl(ctx context.Context, shortKey string) (*url.CreateUrlResponse, error) {
	val, err := a.Client.HGet(ctx, shortKey, shortKey).Result()

	if err != nil {
		return nil, err
	}

	var response url.CreateUrlResponse
	err = json.Unmarshal([]byte(val), &response)
	if err != nil {
		log.Fatalf("Error while unmarshaling %s", err)
	}

	return &response, nil
}

func (a *Adapter) SetUrl(ctx context.Context, shortKey string, url *url.CreateUrlResponse) error {
	jsonData, err := json.Marshal(url)
	if err != nil {
		log.Fatal(err)
	}
	err = a.Client.HSet(ctx, shortKey, shortKey, jsonData).Err()
	if err != nil {
		return err
	}
	return nil
}
