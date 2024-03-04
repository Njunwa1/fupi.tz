package utils

import (
	"math/rand"
	"time"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

// GenerateKey generates a random key
func GenerateKey(outputLength int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, outputLength)
	for i := 0; i < outputLength; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
