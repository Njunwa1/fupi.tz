package grpc

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/redirect"
	"google.golang.org/grpc/metadata"
	"log/slog"
)

func (a Adapter) Redirect(ctx context.Context, request *redirect.RedirectRequest) (*redirect.RedirectResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	slog.Info("Getting Url and redirecting", "request", request, "metadata", md)
	res, err := a.api.Redirect(ctx, request, md)
	if err != nil {
		slog.Error("Failed Redirect", "err", err)
		return &redirect.RedirectResponse{}, err
	}
	return res, nil
}
