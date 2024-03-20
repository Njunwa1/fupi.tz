package shortener

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
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
		log.Fatal("Failed to dial shortener service: ", err)
		return nil
	}
	client := url.NewUrlClient(conn)
	return &Adapter{shortener: client}
}

func (a *Adapter) GetUrlByShortKey(ctx context.Context, shortKey string) (*url.UrlResponse, error) {
	urlObject, err := a.shortener.GetUrlByKey(ctx, &url.UrlByKeyRequest{ShortUrl: shortKey})
	if err != nil {
		return nil, err
	}
	return urlObject, nil
}
