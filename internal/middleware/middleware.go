package middleware

import (
	"cc-luhn-validator/internal/models"
	"encoding/json"
	"net/http"
)

func ValidateJSONHeaderRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			response := models.ErrorResponse{
				Message: "Invalid content type found in request header",
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		next(w, r)
	}
}
