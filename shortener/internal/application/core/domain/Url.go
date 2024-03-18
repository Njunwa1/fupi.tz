package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UrlType struct {
	Name string
}

type Url struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	WebUrl      string
	AndroidUrl  string
	IOSUrl      string
	Short       string
	UserID      *primitive.ObjectID `json:"user_id" bson:"user_id,omitempty"`
	UrlType     UrlType
	CustomAlias string
	Password    string
	ExpiryAt    time.Time
	QrCodeUrl   string
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
	}
}
