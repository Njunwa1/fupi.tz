package grpc

import (
	"github.com/Njunwa1/fupi.tz/plan/internal/ports"
	"google.golang.org/grpc"
)

type Adapter struct {
	port   int
	api    ports.APIPort
	server *grpc.Server
}
