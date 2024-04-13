package utils

import (
	"fmt"
	"github.com/skip2/go-qrcode"
)

type SimpleQRCode struct {
	Content string
	Size    int
}

func (code *SimpleQRCode) Generate() ([]byte, error) {
	qrCode, err := qrcode.Encode(code.Content, qrcode.Medium, code.Size)
	if err != nil {
		return []byte{}, fmt.Errorf("could not generate a QR code: %v", err)
	}
	return qrCode, nil
}
