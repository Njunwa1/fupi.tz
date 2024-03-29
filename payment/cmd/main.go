package main

import (
	"github.com/Njunwa1/fupi.tz/payment/config"
	"github.com/Njunwa1/fupi.tz/payment/internal/adapters/db"
	"github.com/Njunwa1/fupi.tz/payment/internal/adapters/grpc"
	"github.com/Njunwa1/fupi.tz/payment/internal/application/core"
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
		os.Exit(1)
	}

	application := core.NewApplication(dbAdapter)
	if err != nil {
		slog.Info("Error while creating Payments", "err", err)
	}

	grpcAdapter := grpc.NewAdapter(config.GetApplicationPort(), application)
	grpcAdapter.Run()
}