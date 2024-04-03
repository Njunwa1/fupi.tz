package api

import (
	"context"
	"fmt"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/qrcode"
	"github.com/Njunwa1/fupi.tz/shortener/internal/application/core/validation"
	"github.com/Njunwa1/fupi.tz/shortener/internal/ports"
	"github.com/Njunwa1/fupi.tz/shortener/internal/utils"
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"log/slog"
	"time"
)

type Application struct {
	db     ports.DBPort
	keygen ports.KeyGenPort
}

func NewApplication(db ports.DBPort, keygen ports.KeyGenPort) *Application {
	return &Application{db: db, keygen: keygen}
}

func (a *Application) CreateShortUrl(ctx context.Context, request *url.UrlRequest) (*url.UrlResponse, error) {
	err := validation.ValidateUrlCreation(request)
	if err != nil {
		slog.Error("Error while validating Url request")
		return &url.UrlResponse{}, err
	}
	newUrl, err := domain.CreateNewUrl(ctx, request)
	if err != nil {
		slog.Error("Error while Creating new url")
		return &url.UrlResponse{}, err
	}
	//shortURL will equal CustomAlias if it is not empty
	if newUrl.CustomAlias == "" {
		newUrl.Short, err = a.keygen.GenerateShortUrlKey(ctx)
		if err != nil {
			log.Println("Failed to generate short url: ", err)
			return &url.UrlResponse{}, validation.KeyGenerationError(err)
		}
		log.Println("Generated short url", newUrl.Short)
	} else {
		newUrl.Short = newUrl.CustomAlias
	}

	//Generate Qrcode
	qrCode := qrcode.SimpleQRCode{Content: newUrl.Short, Size: 256}
	qrCodeUrl, err := qrCode.Generate(newUrl.Short)
	newUrl.QrCodeUrl = qrCodeUrl

	//Save to DB URL
	insertedUrl, err := a.db.SaveUrl(ctx, *newUrl)
	if err != nil {
		log.Println("Failed to save url to database: ", err)
		return &url.UrlResponse{}, err
	}
	return &url.UrlResponse{
		Id:          insertedUrl.Id.Hex(),
		Type:        insertedUrl.UrlType.Name,
		WebUrl:      insertedUrl.WebUrl,
		IosUrl:      insertedUrl.IOSUrl,
		AndroidUrl:  insertedUrl.AndroidUrl,
		Short:       insertedUrl.Short,
		ExpiryAt:    insertedUrl.ExpiryAt.Format(time.RFC3339),
		CustomAlias: insertedUrl.CustomAlias,
		Password:    insertedUrl.Password,
		QrcodeUrl:   insertedUrl.QrCodeUrl,
	}, nil
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
	userHex, ok := ctx.Value(utils.UserIDKey{}).(string)
	if !ok {
		slog.Error("Failed to get user id from context")
		return &url.UserUrlsResponse{}, fmt.Errorf("failed to get user id from context")
	}
	userObjectID, err := primitive.ObjectIDFromHex(userHex)
	if err != nil {
		slog.Error("Error while converting Object ID", "err", err)
		return &url.UserUrlsResponse{}, err
	}

	urls, err := a.db.GetAllUserUrls(ctx, &userObjectID)
	if err != nil {
		slog.Error("Error while converting Object ID", "err", err)
		return &url.UserUrlsResponse{}, err
	}
	var userUrlsResponse []*url.UrlResponse

	for _, urlData := range urls {
		urlDataId := urlData.Id.Hex()
		expiryAt := urlData.ExpiryAt.Format(time.RFC3339)
		response := url.UrlResponse{
			Id:          urlDataId,
			Short:       urlData.Short,
			WebUrl:      urlData.WebUrl,
			IosUrl:      urlData.IOSUrl,
			AndroidUrl:  urlData.AndroidUrl,
			Type:        urlData.UrlType.Name,
			CustomAlias: urlData.CustomAlias,
			Password:    urlData.Password,
			ExpiryAt:    expiryAt,
			QrcodeUrl:   urlData.QrCodeUrl,
		}
		userUrlsResponse = append(userUrlsResponse, &response)
	}
	return &url.UserUrlsResponse{Urls: userUrlsResponse}, nil
}
