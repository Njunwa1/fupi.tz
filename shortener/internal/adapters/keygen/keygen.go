package keygen

import (
	"context"
	"fupi.tz/proto/golang/keygen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	keygen keygen.KeygenClient
}

func NewAdapter(keygenServiceUrl string) *Adapter {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(keygenServiceUrl, opts...)
	if err != nil {
		panic(err)
	}
	client := keygen.NewKeygenClient(conn)
	return &Adapter{keygen: client}
}

func (a *Adapter) GenerateShortUrlKey(ctx context.Context) (string, error) {
	resp, err := a.keygen.Generate(ctx, &keygen.GenerateKeyRequest{})
	if err != nil {
		return "", err
	}
	return resp.ShortUrl, nil
}
