package domain

import (
	"context"
	"fmt"
	"github.com/Njunwa1/fupi.tz/shortener/internal/utils"
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

type UrlType struct {
	Name string
}

type Url struct {
	Id          primitive.ObjectID  `json:"id" bson:"_id,omitempty"`
	WebUrl      string              `json:"web_url" bson:"web_url"`
	AndroidUrl  string              `json:"android_url" bson:"android_url"`
	IOSUrl      string              `json:"ios_url" bson:"ios_url"`
	Domain      string              `json:"domain" bson:"domain"`
	Short       string              `json:"short" bson:"short"`
	UserID      *primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	UrlType     UrlType             `json:"url_type" bson:"url_type"`
	CustomAlias string              `json:"custom_alias" bson:"custom_alias"`
	Password    string              `json:"password" bson:"password"`
	ExpiryAt    time.Time           `json:"expiry_at" bson:"expiry_at"`
	QrCodeUrl   string              `json:"qrcode_url" bson:"qrcode_url"`
	CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
}

func NewUrl(urlType UrlType, customAlias, password, qrCodeUrl, webUrl, iOSUrl, androidUrl, domain string, userID primitive.ObjectID, expiryAt time.Time) *Url {
	return &Url{
		Short:       "",
		WebUrl:      webUrl,
		AndroidUrl:  androidUrl,
		IOSUrl:      iOSUrl,
		UserID:      &userID,
		UrlType:     urlType,
		CustomAlias: customAlias,
		Password:    password,
		ExpiryAt:    expiryAt,
		QrCodeUrl:   qrCodeUrl,
		Domain:      domain,
		CreatedAt:   time.Now(),
	}
}

func CreateNewUrl(ctx context.Context, request *url.UrlRequest) (*Url, error) {
	userID, ok := ctx.Value(utils.UserIDKey{}).(string)
	if !ok {
		slog.Error("Failed to get user id from context")
		return &Url{}, fmt.Errorf("failed to get user id from context")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	expiryAt, err := time.Parse(time.RFC3339, request.ExpiryAt)
	if err != nil {
		slog.Error("Error while parsing expiry date", "err", err)
		return &Url{}, err
	}
	userIdHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		slog.Error("Error while converting Object ID", "err", err)
		return &Url{}, err
	}

	return NewUrl(
		UrlType{Name: request.Type},
		request.CustomAlias,
		string(hashedPassword),
		request.QrcodeUrl,
		request.WebUrl,
		request.IosUrl,
		request.AndroidUrl,
		"fupi.tz", //replace with request.domain
		userIdHex,
		expiryAt,
	), nil
}
