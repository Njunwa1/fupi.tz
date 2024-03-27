package payment

import (
	"context"
	"github.com/Njunwa1/fupi.tz/subscription/internal/application/domain"
	"github.com/Njunwa1/fupitz-proto/golang/payment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment payment.PaymentClient
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := payment.NewPaymentClient(conn)
	return &Adapter{payment: client}, nil
}

func (a *Adapter) CreatePayment(ctx context.Context, subscription domain.Subscription) (*payment.PaymentResponse, error) {
	res, err := a.payment.CreatePayment(ctx, &payment.PaymentRequest{
		SubscriptionId: subscription.ID.Hex(),
		Amount:         subscription.TotalPrice(),
		Currency:       "USD",
		PaymentMethod:  "PayPal",
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
