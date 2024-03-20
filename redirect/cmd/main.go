package main

import (
	"github.com/Njunwa1/fupi.tz/redirect/config"
	"github.com/Njunwa1/fupi.tz/redirect/internal/adapters/grpc"
	"github.com/Njunwa1/fupi.tz/redirect/internal/adapters/rabbitmq"
	"github.com/Njunwa1/fupi.tz/redirect/internal/adapters/redis"
	"github.com/Njunwa1/fupi.tz/redirect/internal/adapters/shortener"
	"github.com/Njunwa1/fupi.tz/redirect/internal/application/core/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
)

func main() {

	exp, err := exporter(config.GetJaegerURL())
	if err != nil {
		log.Fatal(err)
	}

	tp, err := tracerProvider(exp)
	if err != nil {
		log.Fatal(err)
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}))

	shortenerAdapter := shortener.NewAdapter(config.GetShortenerURL())
	redisAdapter := redis.NewAdapter(config.GetRedisURL())
	rabbitmqAdapter := rabbitmq.NewAdapter(config.GetRabbitMQURL())
	application := api.NewApplication(shortenerAdapter, redisAdapter, rabbitmqAdapter)

	//start the server
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
