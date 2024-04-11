package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/domain"
)

type DBPort interface {
	Save(ctx context.Context, qrCode *domain.QRCode) (*domain.QRCode, error)
	Get(ctx context.Context, id string) (*domain.QRCode, error)
	GetAll(ctx context.Context) ([]*domain.QRCode, error)
	Update(ctx context.Context, qrCode *domain.QRCode) (*domain.QRCode, error)
}
