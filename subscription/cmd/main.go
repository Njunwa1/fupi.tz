package main

import (
	"github.com/Njunwa1/fupi.tz/subscription/config"
	"github.com/Njunwa1/fupi.tz/subscription/internal/adapters/db"
	"github.com/Njunwa1/fupi.tz/subscription/internal/adapters/grpc"
	"github.com/Njunwa1/fupi.tz/subscription/internal/adapters/payment"
	"github.com/Njunwa1/fupi.tz/subscription/internal/adapters/plan"
	"github.com/Njunwa1/fupi.tz/subscription/internal/application/core"
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
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil {
		slog.Info("Error while creating Payment adapter", "err", err)
	}
	planAdapter, err := plan.NewAdapter(config.GetPlanServiceUrl())
	if err != nil {
		slog.Info("Error while creating Plan adapter", "err", err)
	}
	application := core.NewApplication(dbAdapter, paymentAdapter, planAdapter)
	if err != nil {
		slog.Info("Error while creating ", "err", err)
	}

	grpcAdapter := grpc.NewAdapter(config.GetApplicationPort(), application)
	grpcAdapter.Run()
}
