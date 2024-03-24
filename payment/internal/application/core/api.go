package core

import (
	"context"
	"github.com/Njunwa1/fupi.tz/payment/internal/application/domain"
	"github.com/Njunwa1/fupi.tz/payment/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/payment"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) SavePayment(ctx context.Context, request *payment.PaymentRequest) (*payment.PaymentResponse, error) {
	subscriptionId, _ := primitive.ObjectIDFromHex(request.SubscriptionId)
	newPayment := domain.NewPayment(subscriptionId, request.GetCurrency(), request.GetPaymentMethod(), request.GetAmount())
	result, err := a.db.CreatePayment(ctx, *newPayment)
	if err != nil {
		return nil, err
	}
	return &payment.PaymentResponse{
		Id:             result.ID.Hex(),
		SubscriptionId: result.SubscriptionID.Hex(),
		Amount:         result.Amount,
		Currency:       result.Currency,
		PaymentMethod:  result.PaymentMethod,
	}, nil
}
func (a *Application) GetPaymentsByUserID(ctx context.Context, request *payment.UserPaymentsRequest) ([]*payment.PaymentResponse, error) {
	payments, err := a.db.GetUserPayments(ctx, request.GetUserId())
	if err != nil {
		return nil, err
	}
	var response []*payment.PaymentResponse
	for _, p := range payments {
		res := payment.PaymentResponse{
			Id:             p.ID.Hex(),
			SubscriptionId: p.SubscriptionID.Hex(),
			Amount:         p.Amount,
			Currency:       p.Currency,
			PaymentMethod:  p.PaymentMethod,
		}
		response = append(response, &res)
	}
	return response, nil
}
func (a *Application) GetPayments(ctx context.Context) ([]*payment.PaymentResponse, error) {
	payments, err := a.db.GetAllPayments(ctx)
	if err != nil {
		return nil, err
	}
	var response []*payment.PaymentResponse
	for _, p := range payments {
		res := payment.PaymentResponse{
			Id:             p.ID.Hex(),
			SubscriptionId: p.SubscriptionID.Hex(),
			Amount:         p.Amount,
			Currency:       p.Currency,
			PaymentMethod:  p.PaymentMethod,
		}
		response = append(response, &res)
	}
	return response, nil
}
