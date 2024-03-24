package grpc

import (
	"fmt"
	"github.com/Njunwa1/fupi.tz/plan/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/plan"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
)

type Adapter struct {
	port   int
	api    ports.APIPort
	server *grpc.Server
	plan.UnimplementedPlanServer
}

func NewAdapter(port int, api ports.APIPort) *Adapter {
	return &Adapter{
		port: port,
		api:  api,
	}
}

func (a Adapter) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		slog.Error("Error while connecting to server", "port", a.port)
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	)
	a.server = grpcServer

	err = grpcServer.Serve(listener)
	if err != nil {
		slog.Error("Could not serve GRPC on listener", "err", err)
		os.Exit(1)
	}
}
