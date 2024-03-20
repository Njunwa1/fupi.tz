package keygen

import (
	"context"
	"fmt"
	"github.com/Njunwa1/fupitz-proto/golang/keygen"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Adapter struct {
	keygen keygen.KeygenClient
}

func NewAdapter(keygenServiceUrl string) *Adapter {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	opts = append(opts, grpc.WithStatsHandler(otelgrpc.NewClientHandler()))
	conn, err := grpc.Dial(keygenServiceUrl, opts...)
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	client := keygen.NewKeygenClient(conn)
	return &Adapter{keygen: client}
}

func (a *Adapter) GenerateShortUrlKey(ctx context.Context) (string, error) {
	resp, err := a.keygen.Generate(ctx, &keygen.GenerateKeyRequest{})
	fmt.Println(resp)
	if err != nil {
		log.Fatalf("failed to generate key: %v", err)
		return "", err
	}
	return resp.ShortUrl, nil
}
