package grpc

import (
	"fmt"
	"github.com/Njunwa1/fupi.tz/redirect/config"
	"github.com/Njunwa1/fupi.tz/redirect/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/redirect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"log/slog"
	"net"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	redirect.UnimplementedRedirectServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
	grpcServer := grpc.NewServer()
	a.server = grpcServer
	redirect.RegisterRedirectServer(grpcServer, a)

	if config.GetEnv() == "development" {
		reflection.Register(grpcServer)
	}

	slog.Info("starting url service on port  ...", "port", a.port)
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
