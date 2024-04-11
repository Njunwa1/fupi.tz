package core

import (
	"context"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/utils"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/validation"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/qrcode"
	"log/slog"
)

type Application struct {
	db ports.DBPort
	//keygen ports.KeyGenPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a *Application) GenerateQrCode(ctx context.Context, request *qrcode.CreateQRCodeRequest) (*qrcode.QRCodeResponse, error) {
	//Validate request
	err := validation.ValidateQrCodeRequest(request)
	if err != nil {
		slog.Error("Error while validating QRCode request")
		return &qrcode.QRCodeResponse{}, err
	}

	//Generate shortUrl key
	//shortUrl, err := a.keygen.GenerateShortUrlKey(ctx)

	//Generate Qrcode
	qrCode := utils.SimpleQRCode{Content: request.GetShortUrl(), Size: 256}
	fileName := request.GetShortUrl()
	_, err = qrCode.Generate(fileName)
	if err != nil {
		slog.Error("Error while generating QR code", err)
	}

	////Create new QRCode
	//newQrCode := domain.NewQrCode(
	//	request.GetDestinationUrl(),
	//	request.GetShortUrl(),
	//	request.GetBackgroundColor(),
	//	request.GetForegroundColor(),
	//	request.GetLogoUrl(),
	//	request.GetFrameColor(),
	//	request.GetFrameText(),
	//	request.GetBranded(),
	//	//request.GetQrcodeUrl(),
	//	//request.GetUserId(),
	//)

	//Save to DB URL
	//savedQrCode, err = a.db.Save(ctx, newQrCode)

	return &qrcode.QRCodeResponse{
		DestinationUrl:  "",
		ShortUrl:        "",
		BackgroundColor: "",
		ForegroundColor: "",
		LogoPath:        "",
		FrameColor:      "",
		FrameText:       "",
		Branded:         false,
		QrcodeUrl:       "",
		Id:              "",
	}, nil
}
