package handlers

import (
	"cc-luhn-validator/internal/models"
	"cc-luhn-validator/internal/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type ValidationHandler struct {
	validationService service.CardValidationService
}

func NewHandler(validationService service.CardValidationService) *ValidationHandler {
	return &ValidationHandler{
		validationService: validationService,
	}
}

func (h *ValidationHandler) GetValidation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Set response type to JSON format
		w.Header().Set("Content-Type", "application/json")

		var req models.CardValidationRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Println("Bad request: invalid JSON format")

			response := models.ErrorResponse{
				Message: "Invalid JSON format",
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		result, err := h.validationService.ValidateCard(req.CardNumber)

		if err != nil {
			fmt.Printf("Bad request: %s\n", err.Error())

			response := models.ErrorResponse{
				Message: err.Error(),
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := models.CardValidationResponse{
			IsValid:     result.IsValid,
			CardNetwork: result.CardNetwork,
			Source:      result.Source,
		}

		json.NewEncoder(w).Encode(response)
	default:
		fmt.Println("Bad request: method not allowed")

		response := models.ErrorResponse{
			Message: "Method not allowed",
		}

		json.NewEncoder(w).Encode(response)
	}
}
