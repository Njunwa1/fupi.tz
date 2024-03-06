package cmd

import (
	"github.com/njunwa1/fupi.tz/keygen/internal/adapters/db"
	"github.com/njunwa1/fupi.tz/keygen/internal/adapters/grpc"
	"github.com/njunwa1/fupi.tz/keygen/internal/application/core/api"
	"github.com/njunwa1/fupi.tz/shortener/config"
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
