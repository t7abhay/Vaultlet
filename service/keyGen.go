package service

import (
	"crypto"
	crand "crypto/rand"
	"encoding/hex"
	"io"
	"math/rand"
	"time"
	"unicode"
)

func ApiKeyGenerator() (string, error) {
	randomBytes := make([]byte, 256)

	channel := make(chan string)

	_, err := io.ReadFull(crand.Reader, randomBytes)
	if err != nil {
		return "", err
	}

	hasher := crypto.SHA512.New()
	hasher.Write(randomBytes)
	digest := hasher.Sum(nil)

	preApiKey := hex.EncodeToString(digest)

	go characterChanger(preApiKey, channel)
	apiKey := <-channel

	return apiKey, nil
}

func characterChanger(preApiKey string, channel chan string) {
	apiRune := []rune(preApiKey)
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomItterLen := rand.Intn(len(apiRune)) - 1

	for idx := range randomItterLen {
		if unicode.IsLower(apiRune[idx]) {
			apiRune[idx] = unicode.ToUpper(apiRune[idx])

		}
	}

	symbols := []rune{'-', '_'}
	insertCount := rand.Intn(3) + 2

	for range insertCount {
		symbol := symbols[rand.Intn(len(symbols))]
		position := rand.Intn(len(apiRune))
		apiRune = append(apiRune[:position], append([]rune{symbol}, apiRune[position:]...)...)
	}

	channel <- string(apiRune)

}
