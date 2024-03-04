package cmd

import (
	"fupi.tz/shortener/config"
	"fupi.tz/shortener/internal/adapters/db"
	"fupi.tz/shortener/internal/adapters/grpc"
	"fupi.tz/shortener/internal/adapters/keygen"
	"fupi.tz/shortener/internal/application/core/api"
	"log"
)

func main() {

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	keyGenAdapter := keygen.NewAdapter(config.GetKeyGenServiceUrl())

	application := api.NewApplication(dbAdapter, keyGenAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
