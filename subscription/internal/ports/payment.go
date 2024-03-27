package ports

import (
	"context"
	"github.com/Njunwa1/fupi.tz/subscription/internal/application/domain"
	"github.com/Njunwa1/fupitz-proto/golang/payment"
)

type PaymentPort interface {
	CreatePayment(ctx context.Context, subscription domain.Subscription) (*payment.PaymentResponse, error)
}
