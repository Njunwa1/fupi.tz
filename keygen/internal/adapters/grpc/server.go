package grpc

import (
	"fmt"
	"github.com/Njunwa1/fupi.tz-proto/golang/keygen"
	"github.com/Njunwa1/keygen/config"
	"github.com/Njunwa1/keygen/internal/ports"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	keygen.UnimplementedKeygenServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", a.port))
	if err != nil {
		log.Fatal("Failed to listen on port", a.port, "error", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
	)

	a.server = grpcServer

	keygen.RegisterKeygenServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("Failed to serve grpc on port", a.port)
	}
}
