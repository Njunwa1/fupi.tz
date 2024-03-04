package grpc

import (
	"fmt"
	"fupi.tz/keygen/internal/ports"
	"fupi.tz/proto/golang/keygen"
	"fupi.tz/shortener/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		panic(err)
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
		panic(err)
	}
}
