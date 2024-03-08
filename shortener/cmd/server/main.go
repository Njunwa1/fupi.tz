package main

import (
	"context"
	"github.com/Njunwa1/fupi.tz/shortener/config"
	"github.com/Njunwa1/fupi.tz/shortener/internal/adapters/db"
	"github.com/Njunwa1/fupi.tz/shortener/internal/adapters/grpc"
	"github.com/Njunwa1/fupi.tz/shortener/internal/adapters/keygen"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/api"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var mongoClient *mongo.Client

func main() {

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	defer func() {
		if err = dbAdapter.Client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from database: %s", err)
		}
	}()

	keyGenAdapter := keygen.NewAdapter(config.GetKeyGenServiceUrl())

	application := api.NewApplication(dbAdapter, keyGenAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
