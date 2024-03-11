package main

import (
	"github.com/Njunwa1/fupi.tz/redirect/config"
	"github.com/Njunwa1/fupi.tz/redirect/internal/adapters/grpc"
	"github.com/Njunwa1/fupi.tz/redirect/internal/adapters/rabbitmq"
	"github.com/Njunwa1/fupi.tz/redirect/internal/adapters/redis"
	"github.com/Njunwa1/fupi.tz/redirect/internal/adapters/shortener"
	"github.com/Njunwa1/fupi.tz/redirect/internal/application/core/api"
)

func main() {

	shortenerAdapter := shortener.NewAdapter(config.GetShortenerURL())
	redisAdapter := redis.NewAdapter(config.GetRedisURL())
	rabbitmqAdapter := rabbitmq.NewAdapter(config.GetRabbitMQURL())
	application := api.NewApplication(shortenerAdapter, redisAdapter, rabbitmqAdapter)

	//start the server
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
