package ports

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/payment"
)

type APIPort interface {
	SavePayment(ctx context.Context, request *payment.PaymentRequest) (*payment.PaymentResponse, error)
	GetPaymentsByUserID(ctx context.Context, request *payment.UserPaymentsRequest) ([]*payment.PaymentResponse, error)
	GetPayments(ctx context.Context) ([]*payment.PaymentResponse, error)
}
