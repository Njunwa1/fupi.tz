package grpc

import (
	"fmt"
	"github.com/Njunwa1/fupi.tz-proto/golang/clicks"
	"github.com/Njunwa1/fupi.tz/redirect/internal/ports"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	clicks.UnimplementedUrlClicksServer
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
	clicks.RegisterUrlClicksServer(grpcServer, a)

	log.Printf("starting url service on port %d ...", a.port)
	if err := grpcServer.Serve(conn); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
