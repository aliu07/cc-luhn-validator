package server

import (
	"cc-luhn-validator/internal/handlers"
	"cc-luhn-validator/internal/middleware"
	"log"
	"net/http"
)

type CardValidationServer interface {
	SetupRoutes()
}

type cardValidationServer struct {
	ID           string
	Handler      *handlers.ValidationHandler
	ServerLogger *log.Logger
	CacheLogger  *log.Logger
	ErrorLogger  *log.Logger
}

func NewValidationServer(id string, h *handlers.ValidationHandler, serverLogger *log.Logger, cacheLogger *log.Logger, errorLogger *log.Logger) CardValidationServer {
	return &cardValidationServer{
		ID:           id,
		Handler:      h,
		ServerLogger: serverLogger,
		CacheLogger:  cacheLogger,
		ErrorLogger:  errorLogger,
	}
}

func (s *cardValidationServer) SetupRoutes() {
	// Wrap handler with middleware
	http.HandleFunc("/validate", middleware.ValidateJSONHeaderRequest(s.Handler.GetValidation))
}
