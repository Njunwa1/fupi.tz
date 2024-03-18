package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/clicks"
	"google.golang.org/grpc/metadata"
)

type APIPort interface {
	CreateUrlClick(
		ctx context.Context,
		request *clicks.UrlClickRequest,
		md metadata.MD) (*clicks.UrlClickResponse, error)
}
