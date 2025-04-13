package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetValidation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// Set response type to JSON format
		w.Header().Set("Content-Type", "application/json")

		var req CardValidationRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Printf("Bad request: invalid JSON format, returning %d...\n", http.StatusBadRequest)

			response := CardValidationResponse{
				IsValid: false,
				Message: "Invalid JSON format",
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		if req.CardNumber == "" {
			fmt.Printf("Bad request: missing card number, returning %d...\n", http.StatusBadRequest)

			response := CardValidationResponse{
				IsValid: false,
				Message: "Missing card number",
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		var cardNumbers []int

		for _, char := range req.CardNumber {
			digit, err := strconv.Atoi(string(char))

			if err != nil {
				fmt.Printf("Bad request: invalid card number format, returning %d...\n", http.StatusBadRequest)

				response := CardValidationResponse{
					IsValid: false,
					Message: "Invalid card number format",
				}

				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(response)
				return
			}

			cardNumbers = append(cardNumbers, digit)
		}

		isValid := ValidateCard(cardNumbers)

		response := CardValidationResponse{
			IsValid: isValid,
		}

		json.NewEncoder(w).Encode(response)
	default:
		fmt.Printf("Bad request: method not allowed, returning %d...\n", http.StatusMethodNotAllowed)

		response := CardValidationResponse{
			IsValid: false,
			Message: "Method not allowed",
		}

		json.NewEncoder(w).Encode(response)
	}
}
