package service

import (
	"errors"
)

var ErrInvalidApiKey = errors.New("Invalid Api Key ")

func ValidateHashedApiKey(storedHash string, apiKey string) (bool, error) {

	hashToCompare := HashApiKey(apiKey)

	if storedHash != hashToCompare {

		return false, ErrInvalidApiKey
	}

	return true, nil

}
