package main

import (
	"context"
	"github.com/Njunwa1/keygen/config"
	"github.com/Njunwa1/keygen/internal/adapters/db"
	"github.com/Njunwa1/keygen/internal/adapters/grpc"
	"github.com/Njunwa1/keygen/internal/application/core/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	"log/slog"
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
