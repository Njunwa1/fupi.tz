package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/redirect"
)
import "google.golang.org/grpc/metadata"

type RabbitMQPort interface {
	PublishClickEvent(ctx context.Context, urlClick *redirect.RedirectRequest, md metadata.MD) error
}
