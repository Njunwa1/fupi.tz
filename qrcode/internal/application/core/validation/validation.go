package validation

import (
	"github.com/Njunwa1/fupitz-proto/golang/qrcode"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateQrCodeRequest(request *qrcode.CreateQRCodeRequest) error {
	var validationErrors []*errdetails.BadRequest_FieldViolation
	if request.DestinationUrl == "" {
		_ = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "destination_url",
			Description: "Destination URL cannot be empty",
		})
	}
	if len(validationErrors) > 0 {
		stat := status.New(codes.InvalidArgument, "invalid QRCode request")
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return s.Err()
	}

	return nil
}
