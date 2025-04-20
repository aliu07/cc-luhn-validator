package main

import (
	"cc-luhn-validator/internal/cache"
	"cc-luhn-validator/internal/handlers"
	"cc-luhn-validator/internal/server"
	"cc-luhn-validator/internal/service"
	"cc-luhn-validator/internal/utils"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("Setting up server components...")
	cardValidator := utils.NewCardValidator()
	memoryCache := cache.NewLRUMemCache(5)
	cacheTTL := 5 * time.Minute
	service := service.NewCardValidationService(cardValidator, memoryCache, cacheTTL)
	handler := handlers.NewHandler(service)

	fmt.Println("Setting up server...")
	server := server.NewValidationServer("server1", handler)

	fmt.Println("Registering server paths to server mux...")
	server.SetupRoutes()

	// Specify IP address before colon to tell server to listen on specific IP addresses.
	fmt.Println("Listening and ready to serve...")
	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server shutting down...")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
