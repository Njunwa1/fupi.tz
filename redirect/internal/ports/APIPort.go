package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz-proto/golang/clicks"
	"github.com/Njunwa1/fupi.tz-proto/golang/url"
	"google.golang.org/grpc/metadata"
)

type APIPort interface {
	CreateUrlClick(
		ctx context.Context,
		request *clicks.UrlClickRequest,
		md metadata.MD) (*url.CreateUrlResponse, error)
}
