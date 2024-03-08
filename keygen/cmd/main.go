package main

import (
	"context"
	"github.com/Njunwa1/keygen/config"
	"github.com/Njunwa1/keygen/internal/adapters/db"
	"github.com/Njunwa1/keygen/internal/adapters/grpc"
	"github.com/Njunwa1/keygen/internal/application/core/api"
	"log"
	"log/slog"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())

	if err != nil {
		slog.Error("Failed to connect to the database", "error", err)
	}

	defer func() {
		if err = dbAdapter.Client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from database: %s", err)
		}
	}()

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
