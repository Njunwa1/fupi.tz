package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/redirect"
	"google.golang.org/grpc/metadata"
)

type APIPort interface {
	Redirect(
		ctx context.Context,
		request *redirect.RedirectRequest,
		md metadata.MD) (*redirect.RedirectResponse, error)
}
