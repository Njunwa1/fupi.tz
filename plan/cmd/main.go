package main

import (
	"context"
	"github.com/Njunwa1/fupi.tz/plan/config"
	"github.com/Njunwa1/fupi.tz/plan/internal/adapters/db"
	"github.com/Njunwa1/fupi.tz/plan/internal/adapters/grpc"
	"github.com/Njunwa1/fupi.tz/plan/internal/application/core"
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
	_, err = application.CreatePlan(context.Background())
	if err != nil {
		slog.Info("Error while creating plans", "err", err)
	}

	grpcAdapter := grpc.NewAdapter(config.GetApplicationPort(), application)
	grpcAdapter.Run()
}
