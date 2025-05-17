package service

import (
	crand "crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"math/rand"
	"time"
	"unicode"
)

func ApiKeyGenerator() (string, error) {
	randomBytes := make([]byte, 32)

	_, err := io.ReadFull(crand.Reader, randomBytes)
	if err != nil {
		return "", err
	}

	hash := sha256.Sum256(randomBytes)

	preApiKey := hex.EncodeToString(hash[:])

	apiKey := characterChanger(preApiKey)

	return apiKey, nil
}

func characterChanger(preApiKey string) string {
	apiRune := []rune(preApiKey)
	rand.NewSource(time.Now().UnixNano())
	randomItterLen := rand.Intn(len(apiRune)) - 1

	for idx := range randomItterLen {
		if unicode.IsLower(apiRune[idx]) {
			apiRune[idx] = unicode.ToUpper(apiRune[idx])

		}
	}

	symbols := []rune{'-', '_'}
	insertCount := rand.Intn(3) + 2

	for idx := 0; idx < insertCount; idx++ {
		symbol := symbols[rand.Intn(len(symbols))]
		position := rand.Intn(len(apiRune))
		apiRune = append(apiRune[:position], append([]rune{symbol}, apiRune[position:]...)...)
	}

	return string(apiRune)

}
