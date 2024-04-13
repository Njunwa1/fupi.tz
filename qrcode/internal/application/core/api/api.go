package api

import (
	"context"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/core/domain"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/core/utils"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/application/core/validation"
	"github.com/Njunwa1/fupi.tz/qrcode/internal/ports"
	"github.com/Njunwa1/fupitz-proto/golang/qrcode"
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log/slog"
)

type Application struct {
	db        ports.DBPort
	keygen    ports.KeyGenPort
	shortener ports.ShortenerPort
}

func NewApplication(db ports.DBPort, keygen ports.KeyGenPort, shortener ports.ShortenerPort) *Application {
	return &Application{db: db, keygen: keygen, shortener: shortener}
}

func (a *Application) GenerateQrCode(ctx context.Context, request *qrcode.CreateQRCodeRequest) (*qrcode.QRCodeResponse, error) {
	//Validate request
	err := validation.ValidateQrCodeRequest(request)
	if err != nil {
		slog.Error("Error while validating QRCode request")
		return &qrcode.QRCodeResponse{}, err
	}

	//Shorten URL
	_, err = a.shortener.CreateShortUrl(ctx, &url.UrlRequest{
		WebUrl: request.DestinationUrl,
	})
	if err != nil {
		return &qrcode.QRCodeResponse{}, err
	}

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
func (a *Application) GetQrCodes(ctx context.Context, request *qrcode.QRCodesRequest) (*qrcode.QRCodesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) GetQrCode(ctx context.Context, request *qrcode.GetQRCodeRequest) (*qrcode.QRCodeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) UpdateQrCode(ctx context.Context, request *qrcode.UpdateQRCodeRequest) (*qrcode.QRCodeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) UploadQrCodeLogo(server qrcode.QRCode_UploadQRCodeLogoServer) error {
	//TODO implement me
	panic("implement me")
}
