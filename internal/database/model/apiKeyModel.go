package model

import (
	"Vaultlet/service"
	"database/sql"
	"fmt"
	"time"
)

type APIKey struct {
	ID        int            // SERIAL PRIMARY KEY
	UserID    string         // UUID stored as text
	APIKey    string         // the actual key
	Duration  sql.NullString // INTERVAL as string ( e.g. "24 hours")
	CreatedAt time.Time      // TIMESTAMPTZ
}

func SeedDb(db *sql.DB) error {

	const query = `
CREATE TABLE IF NOT EXISTS api_keys (
    id         SERIAL PRIMARY KEY,
    user_id    UUID    NOT NULL,
    api_key    TEXT    NOT NULL,
    duration   INTERVAL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
`
	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("failed to create api_keys table: %w", err)
	}

	fmt.Println("api_keys table is ready")
	return nil
}
func InsertApiKey(db *sql.DB, data APIKey) error {
	const query = `
INSERT INTO api_keys (user_id, api_key, duration, created_at)
VALUES ($1, $2, $3, $4)
`

	_, err := db.Exec(query, data.UserID, data.APIKey, data.Duration.String, data.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert API key: %w", err)
	}

	return nil
}

func ValidateApiKey(db *sql.DB, user_id string, apiKey string) (bool, error) {

	storedKey, err := getStoredHash(db, user_id)

	if err != nil {

		return false, err
	}

	match, err := service.ValidateHashedApiKey(storedKey, apiKey)
	if err != nil {

		return false, fmt.Errorf("validation failed: %w", err)
	}

	return match, nil
}

func getStoredHash(db *sql.DB, user_id string) (string, error) {
	var storedHash string

	query := `SELECT api_key FROM api_keys WHERE user_id = $1`

	res := db.QueryRow(query, user_id).Scan(&storedHash)
	if res != nil {
		if res == sql.ErrNoRows {

			return "", fmt.Errorf("no API key found for user: %s", user_id)
		}
		return "", res
	}

	return storedHash, nil
}

// Function to remove the api key entry for user_id if a creation request is made and the user_id already a apiKey entry in db

func ApiKeyEntryChecker(db *sql.DB, user_id string) (bool, error) {
	var count int

	query := `SELECT COUNT(*) FROM api_keys WHERE user_id = $1`
	err := db.QueryRow(query, user_id).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check existing API key entry: %w", err)
	}

	if count > 0 {

		_, err := apiKeyEntryRemover(db, user_id)
		if err != nil {
			return false, fmt.Errorf("failed to delete existing API key: %w", err)
		}
	}

	return true, nil
}

func apiKeyEntryRemover(db *sql.DB, user_id string) (bool, error) {

	query := `DELETE FROM api_keys WHERE user_id=$1`

	_, err := db.Exec(query, user_id)

	if err != nil {
		return false, fmt.Errorf("failed to delete duplicate entries %w", err)
	}

	return true, nil
}
