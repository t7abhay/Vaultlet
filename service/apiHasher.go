package service

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashApiKey(apiKey string) string {

	apiKeyBytes := []byte(apiKey)

	apiKeyHash := sha256.Sum256(apiKeyBytes)

	hashedApiKey := hex.EncodeToString(apiKeyHash[:])

	return hashedApiKey
}
