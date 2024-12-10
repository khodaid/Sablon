package utils

import (
	"math/rand"
	"strings"
	"time"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	var builder strings.Builder

	for i := 0; i < length; i++ {
		randomIndex := seededRand.Intn(len(charset))
		builder.WriteByte(charset[randomIndex])
	}
	return builder.String()
}

func generateSupplierCode() string {
	code := generateRandomString(8)
	return code
}
