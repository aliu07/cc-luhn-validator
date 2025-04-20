package handlers

import (
	"cc-luhn-validator/internal/models"
	"cc-luhn-validator/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type ValidationHandler struct {
	validationService service.CardValidationService
	errorLogger       *log.Logger
}

func NewHandler(validationService service.CardValidationService, errorLogger *log.Logger) *ValidationHandler {
	return &ValidationHandler{
		validationService: validationService,
		errorLogger:       errorLogger,
	}
}

func (handler *ValidationHandler) GetValidation(writer http.ResponseWriter, reader *http.Request) {
	switch reader.Method {
	case http.MethodPost:
		// Set response type to JSON format
		writer.Header().Set("Content-Type", "application/json")

		var req models.CardValidationRequest

		if err := json.NewDecoder(reader.Body).Decode(&req); err != nil {
			handler.errorLogger.Println("Bad request: invalid JSON format")

			response := models.ErrorResponse{
				Message: "Invalid JSON format",
			}

			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(response)
			return
		}

		result, err := handler.validationService.ValidateCard(req.CardNumber)

		if err != nil {
			handler.errorLogger.Printf("Bad request: %s\n", err.Error())

			response := models.ErrorResponse{
				Message: err.Error(),
			}

			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(response)
			return
		}

		response := models.CardValidationResponse{
			IsValid:     result.IsValid,
			CardNetwork: result.CardNetwork,
			Source:      result.Source,
		}

		json.NewEncoder(writer).Encode(response)
	default:
		handler.errorLogger.Println("Bad request: method not allowed")

		response := models.ErrorResponse{
			Message: "Method not allowed",
		}

		json.NewEncoder(writer).Encode(response)
	}
}
