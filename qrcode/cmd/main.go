package main

import (
	"context"
	"github.com/Njunwa1/fupi.tz/qrcode/config"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/adapters/db"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/adapters/grpc"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/adapters/minio"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/adapters/shortener"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/core/api"
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

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	defer func() {
		if err = dbAdapter.Client.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect from database: %s", err)
		}
	}()

	shortenerAdapter := shortener.NewAdapter(config.GetShortenerServiceUrl())
	minioAdapter := minio.NewAdapter(
		config.GetMinioURL(),
		config.GetMinioAccessKey(),
		config.GetMinioSecretKey(),
		config.GetMinioUseSSL(),
	)

	application := api.NewApplication(dbAdapter, shortenerAdapter, minioAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort())
	grpcAdapter.Run()
}
