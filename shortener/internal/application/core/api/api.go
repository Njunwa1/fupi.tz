package api

import (
	"context"
	"fmt"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/shortener/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"log/slog"
)

type Application struct {
	db     ports.DBPort
	keygen ports.KeyGenPort
}

func NewApplication(db ports.DBPort, keygen ports.KeyGenPort) *Application {
	return &Application{db: db, keygen: keygen}
}

func (a *Application) CreateShortUrl(ctx context.Context, url domain.Url) (domain.Url, error) {
	//shortURL will equal CustomAlias if it is not empty
	if url.CustomAlias == "" {
		log.Println("Sending Original URL to shortening service", url)
		url.Short, _ = a.keygen.GenerateShortUrlKey(ctx)
		log.Println("Generated short url", url.Short)
	} else {
		url.Short = url.CustomAlias
	}
	err := a.db.SaveUrl(ctx, url)
	if err != nil {
		log.Println("Failed to save url to database: ", err)
		return domain.Url{}, err
	}
	return url, nil
}

func (a *Application) GetUrlByShortUrl(ctx context.Context, shortUrl string) (*url.UrlResponse, error) {
	result, err := a.db.GetUrlByShortUrl(ctx, shortUrl)
	fmt.Println("URL from database", result)
	if err != nil {
		log.Println("Failed to get url from database: ", err)
		return &url.UrlResponse{}, err
	}
	res := &url.UrlResponse{
		Id:          result.Id.Hex(),
		Type:        result.UrlType.Name,
		WebUrl:      result.WebUrl,
		IosUrl:      result.IOSUrl,
		AndroidUrl:  result.AndroidUrl,
		Short:       result.Short,
		ExpiryAt:    result.ExpiryAt.Format("2006-01-02"),
		CustomAlias: result.CustomAlias,
		Password:    result.Password,
	}
	return res, nil
}

func (a *Application) GetAllUserUrls(ctx context.Context, request *url.UserUrlsRequest) (*url.UserUrlsResponse, error) {
	userHexID, err := primitive.ObjectIDFromHex(request.GetUserId())
	if err != nil {
		slog.Error("Error while converting Object ID", "err", err)
		return &url.UserUrlsResponse{}, err
	}
	urls, err := a.db.GetAllUserUrls(ctx, &userHexID)
	if err != nil {
		slog.Error("Error while converting Object ID", "err", err)
		return &url.UserUrlsResponse{}, err
	}
	var userUrlsResponse []*url.UrlResponse
	for _, urlData := range urls {
		urlDataId := urlData.Id.Hex()
		expiryAt := urlData.ExpiryAt.Format("2006-01-02")
		response := url.UrlResponse{
			Id:          urlDataId,
			Short:       urlData.Short,
			WebUrl:      urlData.WebUrl,
			IosUrl:      urlData.IOSUrl,
			AndroidUrl:  urlData.AndroidUrl,
			Type:        urlData.UrlType.Name,
			CustomAlias: urlData.CustomAlias,
			ExpiryAt:    expiryAt,
			QrcodeUrl:   urlData.QrCodeUrl,
		}
		_ = append(userUrlsResponse, &response)
	}
	return &url.UserUrlsResponse{Urls: userUrlsResponse}, nil
}
