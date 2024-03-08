package grpc

import (
	"fmt"
	"github.com/Njunwa1/fupi.tz/proto/golang/url"
	"github.com/Njunwa1/fupi.tz/redirect/internal/ports"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	url.UnimplementedUrlServer
}

func NewAdapter(api ports.APIPort, port int) *Adapter {
	return &Adapter{api: api, port: port}
}

func (a Adapter) Run() {
	conn, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		log.Fatalf("failed to listen on port %d, error: %v", a.port, err)
	}
	server := grpc.NewServer()
	a.server = server
	url.RegisterUrlServer(server, a)

	log.Printf("starting url service on port %d ...", a.port)
	if err := server.Serve(conn); err != nil {
		log.Fatalf("failed to serve grpc on port %d", a.port)
	}
}

func (a Adapter) Stop() {
	a.server.Stop()
}
