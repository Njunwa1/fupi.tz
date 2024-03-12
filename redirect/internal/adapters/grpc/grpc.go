package grpc

import (
	"context"
	"github.com/Njunwa1/fupi.tz-proto/golang/clicks"
	"google.golang.org/grpc/metadata"
	"log/slog"
)

func (a Adapter) CreateUrlClick(ctx context.Context, request *clicks.UrlClickRequest) (*clicks.UrlClickResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	slog.Info("Recording url click", "request", request, "metadata", md)
	res, err := a.api.CreateUrlClick(ctx, request, md)
	if err != nil {
		slog.Error("Failed to create url click", "err", err)
		return &clicks.UrlClickResponse{}, err
	}
	return res, nil
}
