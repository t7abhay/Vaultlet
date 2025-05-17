package main

import (
	"Vaultlet/internal/database"
	"Vaultlet/internal/database/model"
	"Vaultlet/internal/handler"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, inboundReq *http.Request) {
		start := time.Now()

		resWriter := &responseWriter{rw: writer, statusCode: http.StatusOK}
		next.ServeHTTP(writer, inboundReq)

		duration := time.Since(start)

		log.Printf("%s %s %d %s", inboundReq.Method, inboundReq.URL.Path, resWriter.statusCode, duration)

	})
}

type responseWriter struct {
	rw         http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.rw.WriteHeader(code)
}

type Handler struct {
	DB *sql.DB
}

func main() {

	loadEnv(".env")
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8011"
	}

	DB, err := database.DbConnection()
	if err != nil {
		panic("Database connection failed")
	}

	if err := model.SeedDb(DB); err != nil {
		log.Fatal("Seeding failed:", err)
	}

	router := http.NewServeMux()
	h := &handler.Handler{
		DB: DB,
	}

	router.HandleFunc("/api/v1/gen-apikey", h.GenerateApiKey)
	router.HandleFunc("/api/v1/validate-apikey", h.ValidateApiKey)

	fmt.Printf("Server started to listen on localhost:%s\n", PORT)
	http.ListenAndServe(":"+PORT, Logger(router))
}

func loadEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal("Error loading .env file")

		panic(1)
	}

	fmt.Println("Env loaded ")
}
