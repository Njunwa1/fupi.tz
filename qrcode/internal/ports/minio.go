package ports

import "context"

type MinioPort interface {
	StoreQRCodeObject(ctx context.Context, bucketName string, objectName string, fileData []byte) (string, error)
}
