package main

import (
	"context"
	"fmt"
	"github.com/Njunwa1/fupi.tz-proto/golang/clicks"
	"github.com/Njunwa1/fupi.tz-proto/golang/url"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	opts := []grpc.DialOption{grpc.WithInsecure()}
	shortenerServiceAddr := "localhost:50051"
	redirectServiceAddr := "localhost:50053"
	mux := runtime.NewServeMux()
	if err := url.RegisterUrlHandlerFromEndpoint(context.Background(), mux, shortenerServiceAddr, opts); err != nil {
		log.Fatalf("failed to register the shortener grpc gateway: %v", err)
	}
	if err := clicks.RegisterUrlClicksHandlerFromEndpoint(context.Background(), mux, redirectServiceAddr, opts); err != nil {
		log.Fatalf("failed to register the redirect grpc gateway: %v", err)
	}

	// start listening to requests from the gateway server
	addr := "0.0.0.0:8080"
	fmt.Println("API gateway server is running on " + addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
