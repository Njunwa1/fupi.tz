package shortener

import (
	"context"
	"github.com/Njunwa1/fupi.tz/redirect/proto/golang/url"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type Adapter struct {
	shortener url.UrlClient
}

func NewAdapter(shortenerServiceUrl string) *Adapter {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(shortenerServiceUrl, opts...)
	if err != nil {
		log.Fatal("Failed to dial shortener service: ", err)
		return nil
	}
	client := url.NewUrlClient(conn)
	return &Adapter{shortener: client}
}

func (s *Adapter) GetUrlByShortKey(ctx context.Context, shortKey string) (*url.CreateUrlResponse, error) {
	urlObject, err := s.shortener.GetUrlByKey(ctx, &url.GetUrlByKeyRequest{ShortUrl: shortKey})
	if err != nil {
		return nil, err
	}
	return urlObject, nil
}
