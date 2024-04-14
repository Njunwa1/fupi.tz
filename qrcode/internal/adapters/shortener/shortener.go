package shortener

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"log/slog"
)

type Adapter struct {
	shortener url.UrlClient
}

func NewAdapter(shortenerServiceUrl string) *Adapter {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	conn, err := grpc.Dial(shortenerServiceUrl, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	client := url.NewUrlClient(conn)
	return &Adapter{shortener: client}
}

func (a *Adapter) CreateShortUrl(ctx context.Context, request *url.UrlRequest) (*url.UrlResponse, error) {
	result, err := a.shortener.CreateShortUrl(ctx, request)
	if err != nil {
		slog.Error("failed to create short url: %v", err)
		return nil, err
	}
	return result, nil
}
