package ports

import "github.com/Njunwa1/fupitz-proto/golang/qrcode"
import "context"

type APIPort interface {
	GetQrCodes(ctx context.Context, request *qrcode.QRCodesRequest) (*qrcode.QRCodesResponse, error)
	GenerateQrCode(ctx context.Context, request *qrcode.CreateQRCodeRequest) (*qrcode.QRCodeResponse, error)
	GetQrCode(ctx context.Context, request *qrcode.GetQRCodeRequest) (*qrcode.QRCodeResponse, error)
	UpdateQrCode(ctx context.Context, request *qrcode.UpdateQRCodeRequest) (*qrcode.QRCodeResponse, error)
	UploadQrCodeLogo(qrcode.QRCode_UploadQRCodeLogoServer) error
}
