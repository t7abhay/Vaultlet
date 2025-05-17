package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func DbConnection() (*sql.DB, error) {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	fmt.Println("Connecting to database ....... ")

	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	DbInstance, err := sql.Open("postgres", psqlconn)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := DbInstance.Ping(); err != nil {

		return nil, fmt.Errorf("failed to ping DB: %w", err)
	}
	fmt.Println("Connected!")
	return DbInstance, nil
}
