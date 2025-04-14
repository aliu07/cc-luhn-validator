package handlers

import (
	"cc-luhn-validator/internal/cache"
	"cc-luhn-validator/internal/models"
	"cc-luhn-validator/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type DataSource int

const (
	Handler DataSource = iota
	Cache
	Server
)

func (d DataSource) String() string {
	switch d {
	case Handler:
		return "handler"
	case Cache:
		return "cache"
	case Server:
		return "server"
	default:
		return "unkown"
	}
}

type ValidationHandler struct {
	validator utils.Validator
	cache     cache.Cache
}

func NewHandler(v utils.Validator, c cache.Cache) *ValidationHandler {
	return &ValidationHandler{
		validator: v,
		cache:     c,
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

			response := models.CardValidationResponse{
				IsValid:     false,
				CardNetwork: utils.Unknown.String(),
				Message:     "Invalid JSON format",
				Source:      Handler.String(),
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Check cache
		if data, exists := h.cache.Get(req.CardNumber); exists {
			response := models.CardValidationResponse{
				IsValid:     data.IsValid,
				CardNetwork: data.CardNetwork,
				Source:      Cache.String(),
			}

			json.NewEncoder(w).Encode(response)
			return
		}

		if req.CardNumber == "" {
			fmt.Println("Bad request: missing card number")

			response := models.CardValidationResponse{
				IsValid:     false,
				CardNetwork: utils.Unknown.String(),
				Message:     "Missing card number",
				Source:      Handler.String(),
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		isValid, err := h.validator.ValidateString(req.CardNumber)
		cardNetwork := utils.GetCardNetwork(req.CardNumber)

		if err != nil {
			fmt.Println("Bad request: invalid character in card number")

			response := models.CardValidationResponse{
				IsValid:     false,
				CardNetwork: utils.Unknown.String(),
				Message:     "Invalid character detected in card number",
				Source:      Handler.String(),
			}

			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}

		// Put in cache
		// TODO: put TTL in config struct
		h.cache.Put(req.CardNumber, isValid, cardNetwork, 5*time.Minute)

		response := models.CardValidationResponse{
			IsValid:     isValid,
			CardNetwork: cardNetwork,
			Source:      Server.String(),
		}

		json.NewEncoder(w).Encode(response)
	default:
		fmt.Println("Bad request: method not allowed")

		response := models.CardValidationResponse{
			IsValid:     false,
			CardNetwork: utils.Unknown.String(),
			Message:     "Method not allowed",
			Source:      Handler.String(),
		}

		json.NewEncoder(w).Encode(response)
	}
}
