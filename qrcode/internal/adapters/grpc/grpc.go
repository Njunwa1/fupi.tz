package grpc

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/qrcode"
	"log/slog"
)

func (a Adapter) GenerateQRCode(ctx context.Context, request *qrcode.CreateQRCodeRequest) (*qrcode.QRCodeResponse, error) {
	slog.Info("Generating QRCode")
	result, err := a.api.GenerateQrCode(ctx, request)
	if err != nil {
		slog.Error("Generate QRCode request failed", "error", err)
		return nil, err
	}
	return result, nil
}
