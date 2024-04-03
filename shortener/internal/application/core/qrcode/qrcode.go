package qrcode

import (
	"bytes"
	"fmt"
	"github.com/skip2/go-qrcode"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

type SimpleQRCode struct {
	Content string
	Size    int
}

func (code *SimpleQRCode) Generate(fileName string) (string, error) {
	qrCode, err := qrcode.Encode(code.Content, qrcode.Medium, code.Size)
	if err != nil {
		return "", fmt.Errorf("could not generate a QR code: %v", err)
	}
	// Create the directory if it doesn't exist
	if err := os.MkdirAll("./public/images", 0755); err != nil {
		return "", fmt.Errorf("error creating directory: %v", err)
	}

	// Create the file path
	qrcodeFilePath := filepath.Join("./public/images", fileName)

	// Open the file for writing
	file, err := os.Create(qrcodeFilePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Decode the QR code byte slice into an image
	qrImage, _, err := image.Decode(bytes.NewReader(qrCode))
	if err != nil {
		return "", fmt.Errorf("error decoding QR code: %v", err)
	}

	// Write the image to the file in PNG format
	if err := png.Encode(file, qrImage); err != nil {
		return "", fmt.Errorf("error encoding PNG: %v", err)
	}
	return qrcodeFilePath, nil
}
