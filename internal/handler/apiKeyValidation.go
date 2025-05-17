package handler

import (
	"Vaultlet/internal/database/model"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Handler struct {
	DB *sql.DB
}
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ValidationRequest struct {
	Userid string `json:"user_id"`
	Apikey string `json:"api_key"`
}

func (h *Handler) ValidateApiKey(writer http.ResponseWriter, request *http.Request) {

	dbInstance := h.DB
	var validationRequest ValidationRequest

	if request.Method != http.MethodPost {
		ApiResponseWriter(writer, http.StatusMethodNotAllowed, nil)
		return
	}

	err := json.NewDecoder(request.Body).Decode(&validationRequest)

	if err != nil {
		ApiResponseWriter(writer, http.StatusBadRequest, &Response{
			Success: false,
			Message: "Invalid request body",
		})
		return
	}

	_, err = model.ValidateApiKey(dbInstance, validationRequest.Userid, validationRequest.Apikey)

	if err != nil {
		ApiResponseWriter(writer, http.StatusUnauthorized, &Response{
			Success: false,
			Message: "Invalid API key",
		})
		return
	}

	response := &Response{
		Success: true,
		Message: "Valid Api Key",
	}
	ApiResponseWriter(writer, http.StatusOK, response)

}
