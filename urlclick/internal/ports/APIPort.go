package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz-proto/golang/clicks"
)

type APIPort interface {
	SaveUrlClick(ctx context.Context, request *clicks.UrlClickRequest) (*clicks.UrlClickResponse, error)
}
