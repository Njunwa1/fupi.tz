package cmd

import (
	"fupi.tz/keygen/internal/adapters/db"
	"fupi.tz/keygen/internal/adapters/grpc"
	"fupi.tz/keygen/internal/application/core/api"
	"fupi.tz/shortener/config"
	"log/slog"
)

func main() {
	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())

	if err != nil {
		slog.Error("Failed to connect to the database", "error", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
