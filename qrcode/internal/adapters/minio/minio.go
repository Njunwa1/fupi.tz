package minio

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
	"log/slog"
)

type Adapter struct {
	endPoint    string
	accessKeyId string
	secretKey   string
	useSSL      bool
	minioClient *minio.Client
}

func NewAdapter(endPoint, accessKeyId, secretKey string, useSSL bool) *Adapter {
	minioClient, err := minio.New(endPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyId, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		slog.Error("Error creating MinIO client:", err)
		log.Fatalln(err)
	}
	return &Adapter{minioClient: minioClient}
}

func (a *Adapter) StoreQRCodeObject(ctx context.Context, bucketName, objectName string, imageData []byte) (string, error) {
	// Store the image to MinIO
	_, err := a.minioClient.PutObject(ctx, bucketName, objectName, bytes.NewReader(imageData), int64(len(imageData)), minio.PutObjectOptions{
		ContentType: "image/png", // Change content type according to your image type
	})
	if err != nil {
		slog.Error("Error uploading image to MinIO:", err)
		return "", err
	}
	imageURL := fmt.Sprintf("http://%s/%s/%s", a.endPoint, bucketName, objectName)
	return imageURL, nil
}
