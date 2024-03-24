package grpc

import (
	"fmt"
	"github.com/Njunwa1/fupi.tz/payment/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/payment"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	payment.UnimplementedPaymentServer
}

func NewAdapter(port int, api ports.APIPort) *Adapter {
	return &Adapter{
		api:  api,
		port: port,
	}
}

func (a Adapter) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		slog.Error("Cannot open connection on this port", "port", a.port)
		os.Exit(1)
	}
	grpcServer := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	a.server = grpcServer
	err = grpcServer.Serve(listener)
	if err != nil {
		slog.Error("Cannot open connection on this port", "port", a.port)
		os.Exit(1)
	}
}
