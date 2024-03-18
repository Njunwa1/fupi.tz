package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	Short       string              `json:"short" bson:"short"`
	UserID      *primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	UrlType     UrlType             `json:"url_type" bson:"url_type"`
	CustomAlias string              `json:"custom_alias" bson:"custom_alias"`
	Password    string              `json:"password" bson:"password"`
	ExpiryAt    time.Time           `json:"expiry_at" bson:"expiry_at"`
	QrCodeUrl   string              `json:"qrcode_url" bson:"qrcode_url"`
	CreatedAt   time.Time           `json:"created_at" bson:"created_at"`
}

func NewUrl(urlType UrlType, customAlias, password, qrCodeUrl, webUrl, iOSUrl, androidUrl string, userID primitive.ObjectID, expiryAt time.Time) *Url {
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
		CreatedAt:   time.Now(),
	}
}
