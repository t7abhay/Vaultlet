package handler

import (
	"Vaultlet/internal/database/model"
	"Vaultlet/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	ApiKey  string `json:"api_key,omitempty"`
}

type RequestData struct {
	UserID   string `json:"user_id"`
	Duration int    `json:"duration"`
}

func (h *Handler) GenerateApiKey(writer http.ResponseWriter, request *http.Request) {

	dbInstance := h.DB

	if request.Method != http.MethodPost {
		ApiResponseWriter(writer, http.StatusMethodNotAllowed, ApiResponse{
			Success: false,
			Message: "Method not allowed",
		})
		return
	}

	body, err := io.ReadAll(request.Body)

	defer request.Body.Close()

	if err != nil {
		ApiResponseWriter(writer, http.StatusBadRequest, ApiResponse{
			Success: false,
			Message: "Failed to read request body",
		})
		return
	}

	var requestObject RequestData
	if err = json.Unmarshal(body, &requestObject); err != nil {
		ApiResponseWriter(writer, http.StatusBadRequest, ApiResponse{
			Success: false,
			Message: "Invalid JSON in request body",
		})
		return
	}

	_, err = model.ApiKeyEntryChecker(dbInstance, requestObject.UserID)

	if err != nil {
		ApiResponseWriter(writer, http.StatusInternalServerError, ApiResponse{
			Success: false,
			Message: "Failed to check and delete existing entry",
		})
		return
	}

	apiKey, err := service.ApiKeyGenerator()

	if err != nil {
		ApiResponseWriter(writer, http.StatusInternalServerError, ApiResponse{
			Success: false,
			Message: "Failed to generate API key",
		})
		return
	}

	hashedApiKey := service.HashApiKey(apiKey)
	apiKeyData := model.APIKey{
		UserID:    requestObject.UserID,
		APIKey:    hashedApiKey,
		Duration:  sql.NullString{String: fmt.Sprintf("%d hours", requestObject.Duration), Valid: true},
		CreatedAt: time.Now(),
	}

	err = model.InsertApiKey(dbInstance, apiKeyData)
	if err != nil {
		ApiResponseWriter(writer, http.StatusInternalServerError, ApiResponse{
			Success: false,
			Message: "Failed to save API key",
		})
		return
	}

	ApiResponseWriter(writer, http.StatusCreated, ApiResponse{
		Success: true,
		Message: "API key created",
		ApiKey:  apiKey,
	})
}

func ApiResponseWriter(writer http.ResponseWriter, status int, data any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	json.NewEncoder(writer).Encode(data)
}
