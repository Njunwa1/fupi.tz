package main

import (
	"context"
	"github.com/Njunwa1/fupi.tz/redirect/config"
	"github.com/Njunwa1/fupi.tz/redirect/internal/adapters/db"
	"github.com/Njunwa1/fupi.tz/redirect/internal/adapters/grpc"
	"github.com/Njunwa1/fupi.tz/redirect/internal/adapters/shortener"
	"github.com/Njunwa1/fupi.tz/redirect/internal/application/core/api"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func main() {
	dbAdapter, _ := db.NewAdapter(config.GetDataSourceURL())

	defer func(Client *mongo.Client, ctx context.Context) {
		err := Client.Disconnect(ctx)
		if err != nil {
			log.Fatal("Failed to disconnect from database: ", err)
		}
	}(dbAdapter.Client, context.Background())

	shortenerAdapter := shortener.NewAdapter(config.GetShortenerURL())

	application := api.NewApplication(dbAdapter, shortenerAdapter)

	//start the server
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
