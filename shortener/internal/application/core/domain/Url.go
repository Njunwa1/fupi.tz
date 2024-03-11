package domain

import (
	"time"
)

type UrlType struct {
	Name string
}

type Url struct {
	Id          string `json:"id" bson:"_id,omitempty"`
	WebUrl      string
	AndroidUrl  string
	IOSUrl      string
	Short       string
	UserID      int64
	UrlType     UrlType
	CustomAlias string
	Password    string
	ExpiryAt    time.Time
	QrCodeUrl   string
}

func NewUrl(userID int64, urlType UrlType, customAlias, password, qrCodeUrl, webUrl, iOSUrl, androidUrl string, expiryAt time.Time) *Url {
	return &Url{
		WebUrl:      webUrl,
		AndroidUrl:  androidUrl,
		IOSUrl:      iOSUrl,
		UserID:      userID,
		UrlType:     urlType,
		CustomAlias: customAlias,
		Password:    password,
		ExpiryAt:    expiryAt,
		QrCodeUrl:   qrCodeUrl,
	}
}
