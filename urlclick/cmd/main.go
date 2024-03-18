package main

import (
	"context"
	"github.com/Njunwa1/fupi.tz/urlclick/config"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/adapters/db"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/adapters/grpc"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/adapters/rabbitmq"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/application/core/api"
	"go.mongodb.org/mongo-driver/mongo"
	"log/slog"
)

func main() {
	dbAdapter := db.NewAdapter(config.GetDataSourceURL())
	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			slog.Error("Failed to disconnect from the database", "err", err)
		}
	}(dbAdapter.Client, context.Background())

	go func() {
		rabbitmqAdapter := rabbitmq.NewAdapter(config.GetRabbitMQURL(), dbAdapter)
		slog.Info("Starting Consuming click events")
		err := rabbitmqAdapter.ConsumeClickEvent()
		if err != nil {
			slog.Error("Failed to consume click event", "err", err)
		}
	}()

	application := api.NewApplication(dbAdapter)

	slog.Info("Starting the server on port", "port", config.GetApplicationPort())
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()

}
