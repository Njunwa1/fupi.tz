package validation

import (
	"github.com/Njunwa1/fupitz-proto/golang/url"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	netUrl "net/url"
	"time"
)

func ValidateUrlCreation(request *url.UrlRequest) error {
	var validationErrors []*errdetails.BadRequest_FieldViolation
	if request.GetWebUrl() != "" && isValidURL(request.GetWebUrl()) {
		_ = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "web_url",
			Description: "This is not a valid url",
		})
	}

	if request.GetAndroidUrl() != "" && isValidURL(request.GetAndroidUrl()) {
		_ = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "android_url",
			Description: "This is not a valid url",
		})
	}
	if request.GetIosUrl() != "" && isValidURL(request.GetIosUrl()) {
		_ = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "ios_url",
			Description: "This is not a valid url",
		})
	}

	expiryDate, err := time.Parse(time.RFC3339, request.GetExpiryAt())
	if err != nil {
		_ = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "expiry_at",
			Description: "This is not a valid date",
		})
	}

	if !expiryDate.After(time.Now()) {
		_ = append(validationErrors, &errdetails.BadRequest_FieldViolation{
			Field:       "expiry_at",
			Description: "Expiry date must be after today",
		})
	}

	if len(validationErrors) > 0 {
		stat := status.New(codes.InvalidArgument, "invalid URL request")
		badRequest := &errdetails.BadRequest{}
		badRequest.FieldViolations = validationErrors
		s, _ := stat.WithDetails(badRequest)
		return s.Err()
	}

	return nil
}

func KeyGenerationError(error error) error {
	st, _ := status.FromError(error)
	fieldErr := &errdetails.BadRequest_FieldViolation{
		Field:       "keygen",
		Description: st.Message(),
	}
	badRequest := errdetails.BadRequest{}
	badRequest.FieldViolations = append(badRequest.FieldViolations, fieldErr)
	shortenerStatus := status.New(codes.InvalidArgument, "ShortUrl keygen error")
	statusDetails, _ := shortenerStatus.WithDetails(&badRequest)
	return statusDetails.Err()
}

func isValidURL(rawURL string) bool {
	_, err := netUrl.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}
	return true
}
