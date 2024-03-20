package main

import (
	"context"
	"github.com/Njunwa1/fupi.tz/auth/config"
	"github.com/Njunwa1/fupi.tz/auth/internal/adapters/db"
	"github.com/Njunwa1/fupi.tz/auth/internal/adapters/grpc"
	"github.com/Njunwa1/fupi.tz/auth/internal/adapters/paseto"
	"github.com/Njunwa1/fupi.tz/auth/internal/application/core/api"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"log"
	"log/slog"
	"os"
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
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}
	defer func() {
		if err = dbAdapter.Client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from database: %s", err)
		}
	}()

	pasetoAdapter, err := paseto.NewPasetoMaker(string(config.GetPasetoSecret()))
	if err != nil {
		slog.Error("Failed to create paseto maker", "err", err)
		os.Exit(1)
	}
	application := api.NewApplication(dbAdapter, pasetoAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
