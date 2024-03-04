package grpc

import (
	"context"
	"fupi.tz/proto/golang/keygen"
	"log/slog"
)

func (a Adapter) generateKey(ctx context.Context, request *keygen.GenerateKeyRequest) (*keygen.GenerateKeyResponse, error) {
	slog.Info("Generate Key request", "request", request)
	shortUrl, err := a.api.GenerateShortUrlKey(ctx)
	if err != nil {
		return nil, err
	}
	return &keygen.GenerateKeyResponse{ShortUrl: shortUrl}, nil
}
