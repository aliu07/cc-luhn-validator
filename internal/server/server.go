package server

import (
	"cc-luhn-validator/internal/handlers"
	"cc-luhn-validator/internal/middleware"
	"net/http"
)

type CardValidationServer interface {
	SetupRoutes()
}

type cardValidationServer struct {
	ID      string
	Handler *handlers.ValidationHandler
}

func NewValidationServer(id string, h *handlers.ValidationHandler) CardValidationServer {
	return &cardValidationServer{
		ID:      id,
		Handler: h,
	}
}

func (s *cardValidationServer) SetupRoutes() {
	// Wrap handler with middleware
	http.HandleFunc("/validate", middleware.ValidateJSONHeaderRequest(s.Handler.GetValidation))
}
