package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/payment/internal/application/domain"
)

type DBPort interface {
	CreatePayment(ctx context.Context, payment domain.Payment) (domain.Payment, error)
	GetUserPayments(ctx context.Context, userId string) ([]domain.Payment, error)
	GetAllPayments(ctx context.Context) ([]domain.Payment, error)
}
