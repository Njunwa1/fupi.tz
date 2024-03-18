package grpc

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/keygen"
	"log/slog"
)

func (a Adapter) Generate(ctx context.Context, request *keygen.GenerateKeyRequest) (*keygen.GenerateKeyResponse, error) {
	slog.Info("Generate Key request", "request", request)
	entry, err := a.api.GenerateShortUrlKey(ctx)
	if err != nil {
		return nil, err
	}
	return &keygen.GenerateKeyResponse{ShortUrl: entry.ShortUrl}, nil
}
