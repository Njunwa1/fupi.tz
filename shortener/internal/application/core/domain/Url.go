package domain

import "time"

type UrlType struct {
	Name string
}

type Url struct {
	Original    string
	Short       string
	UserID      int64
	UrlType     UrlType
	CustomAlias string
	Password    string
	ExpiryAt    time.Time
	QrCodeUrl   string
}

func NewUrl(userID int64, urlType UrlType, customAlias, password, qrCodeUrl, original string, expiryAt time.Time) *Url {
	return &Url{
		Original:    original,
		UserID:      userID,
		UrlType:     urlType,
		CustomAlias: customAlias,
		Password:    password,
		ExpiryAt:    expiryAt,
		QrCodeUrl:   qrCodeUrl,
	}
}
