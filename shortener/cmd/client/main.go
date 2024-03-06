package main

import (
	"context"
	"fmt"
	"github.com/Njunwa1/proto/golang/url"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

var shortenerServiceAddr string

func main() {
	// Set up a connection to the shortener server.
	conn, err := grpc.Dial(shortenerServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	defer conn.Close()
	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	if err = url.RegisterUrlHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the order server: %v", err)
	}
	// start listening to requests from the gateway server
	addr := "0.0.0.0:8080"
	fmt.Println("API gateway server is running on " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
