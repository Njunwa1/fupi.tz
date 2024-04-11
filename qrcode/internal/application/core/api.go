package core

import (
	"context"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/domain"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/utils"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/validation"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/qrcode"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
)

type Application struct {
	db     ports.DBPort
	keygen ports.KeyGenPort
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
	shortUrl, err := a.keygen.GenerateShortUrlKey(ctx)
	request.ShortUrl = shortUrl

	//Generate Qrcode
	qrCode := utils.SimpleQRCode{Content: request.GetShortUrl(), Size: 256}
	fileName := request.GetShortUrl()
	qrcodePath, err := qrCode.Generate(fileName)
	if err != nil {
		slog.Error("Error while generating QR code", err)
		return &qrcode.QRCodeResponse{}, err
	}

	userId, _ := primitive.ObjectIDFromHex(request.GetUserId())

	//Create new QRCode
	newQrCode := domain.NewQrCode(
		request.GetDestinationUrl(),
		request.GetShortUrl(),
		request.GetBackgroundColor(),
		request.GetForegroundColor(),
		request.GetLogoPath(),
		request.GetFrameColor(),
		request.GetFrameText(),
		request.GetBranded(),
		userId,
	)

	//Save to DB URL
	savedQrCode, err := a.db.Save(ctx, newQrCode)
	if err != nil {
		slog.Error("Failed to save QRCode to database", err)
		return &qrcode.QRCodeResponse{}, err
	}

	return &qrcode.QRCodeResponse{
		DestinationUrl:  savedQrCode.DestinationURL,
		ShortUrl:        savedQrCode.ShortURL,
		BackgroundColor: savedQrCode.BackgroundColor,
		ForegroundColor: savedQrCode.ForegroundColor,
		LogoPath:        savedQrCode.Logo,
		FrameColor:      savedQrCode.FrameColor,
		FrameText:       savedQrCode.FrameText,
		Branded:         savedQrCode.Branded,
		QrcodeUrl:       qrcodePath,
		Id:              savedQrCode.ID.Hex(),
	}, nil
}
