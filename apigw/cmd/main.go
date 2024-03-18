package main

import (
	"context"
	"fmt"
	"github.com/Njunwa1/fupitz-proto/golang/clicks"
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"github.com/Njunwa1/fupitz-proto/golang/user"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	shortenerServiceAddr := "localhost:50051"
	redirectServiceAddr := "localhost:50053"
	AggregatorServiceAddr := "localhost:50054"
	AuthServiceAddr := "localhost:50055"
	mux := runtime.NewServeMux()
	if err := url.RegisterUrlHandlerFromEndpoint(context.Background(), mux, shortenerServiceAddr, opts); err != nil {
		slog.Error("failed to register the shortener grpc gateway: %v", "err", err)
	}
	if err := clicks.RegisterUrlClicksHandlerFromEndpoint(context.Background(), mux, redirectServiceAddr, opts); err != nil {
		slog.Error("failed to register the redirect grpc gateway: ", "err", err)
	}
	if err := clicks.RegisterUrlClicksHandlerFromEndpoint(context.Background(), mux, AggregatorServiceAddr, opts); err != nil {
		slog.Error("failed to register the Aggregator grpc gateway:", "err", err)
	}
	if err := user.RegisterUserHandlerFromEndpoint(context.Background(), mux, AuthServiceAddr, opts); err != nil {
		slog.Error("failed to register the auth grpc gateway:", "err", err)
	}

	// start listening to requests from the gateway server
	addr := "0.0.0.0:8080"
	fmt.Println("API gateway server is running on " + addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
