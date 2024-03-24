package grpc

import (
	"context"
	"github.com/Njunwa1/fupitz-proto/golang/payment"
	"log/slog"
)

func (a Adapter) CreatePayment(ctx context.Context, request *payment.PaymentRequest) (*payment.PaymentResponse, error) {
	res, err := a.api.SavePayment(ctx, request)
	if err != nil {
		slog.Error("Error while creating payment", "err", err)
		return nil, err
	}
	return res, nil
}

func (a Adapter) GetUserPayments(ctx context.Context, request *payment.UserPaymentsRequest) ([]*payment.PaymentResponse, error) {
	res, err := a.api.GetPaymentsByUserID(ctx, request)
	if err != nil {
		slog.Error("Error while retrieving payments", "err", err, "user", request.GetUserId())
		return nil, err
	}
	return res, nil
}

func (a Adapter) GetAllPayments(ctx context.Context) ([]*payment.PaymentResponse, error) {
	res, err := a.api.GetPayments(ctx)
	if err != nil {
		slog.Error("Error while retrieving payments", "err", err)
		return nil, err
	}
	return res, nil
}
