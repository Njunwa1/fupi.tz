package grpc

import (
	"fmt"
	"github.com/Njunwa1/fupi.tz-proto/golang/clicks"
	"github.com/Njunwa1/fupi.tz/urlclick/internal/ports"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

type Adapter struct {
	api    ports.APIPort
	server *grpc.Server
	port   int
	clicks.UnimplementedUrlClicksServer
}

func NewAdapter(port int) *Adapter {
	return &Adapter{port: port}
}

func (a Adapter) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		slog.Error("Failed to start the server on port", "port", a.port, "err", err)
	}
	//create the grpc server
	a.server = grpc.NewServer()
	clicks.RegisterUrlClicksServer(a.server, a)
	//start the server
	slog.Info("Starting the server on port", "port", a.port)
	err = a.server.Serve(listener)
	if err != nil {
		slog.Error("Failed to start the server", "err", err)
	}
}
