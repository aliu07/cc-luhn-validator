package internal

import (
	"cc-luhn-validator/internal/models"
	"cc-luhn-validator/internal/validation"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	validator validation.Validator
}

func NewHandler(v validation.Validator) *Handler {
	return &Handler{
		validator: v,
	}
}

func (h *Handler) GetValidation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Set response type to JSON format
		w.Header().Set("Content-Type", "application/json")

		var req models.CardValidationRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Printf("Bad request: invalid JSON format, returning %d...\n", http.StatusBadRequest)

			response := models.CardValidationResponse{
				IsValid: false,
				Message: "Invalid JSON format",
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		if req.CardNumber == "" {
			fmt.Printf("Bad request: missing card number, returning %d...\n", http.StatusBadRequest)

			response := models.CardValidationResponse{
				IsValid: false,
				Message: "Missing card number",
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		isValid, err := h.validator.ValidateString(req.CardNumber)

		if err != nil {
			fmt.Printf("%s, returning %d...", err.Error(), http.StatusBadRequest)

			response := models.CardValidationResponse{
				IsValid: false,
				Message: "Invalid character detected in card number",
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		response := models.CardValidationResponse{
			IsValid: isValid,
		}

		json.NewEncoder(w).Encode(response)
	default:
		fmt.Printf("Bad request: method not allowed, returning %d...\n", http.StatusMethodNotAllowed)

		response := models.CardValidationResponse{
			IsValid: false,
			Message: "Method not allowed",
		}

		json.NewEncoder(w).Encode(response)
	}
}
