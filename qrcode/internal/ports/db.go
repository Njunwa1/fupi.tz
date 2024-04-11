package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/domain"
)

type DBPort interface {
	save(ctx context.Context, qrCode *domain.QRCode) (*domain.QRCode, error)
	get(ctx context.Context, id string) (*domain.QRCode, error)
	getAll(ctx context.Context) ([]*domain.QRCode, error)
	update(ctx context.Context, qrCode *domain.QRCode) (*domain.QRCode, error)
}
