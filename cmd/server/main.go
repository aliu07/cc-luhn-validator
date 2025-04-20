package main

import (
	"cc-luhn-validator/internal/cache"
	"cc-luhn-validator/internal/handlers"
	"cc-luhn-validator/internal/middleware"
	"cc-luhn-validator/internal/utils"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Dummy credit card numbers can be found in stripe testing doc: https://docs.stripe.com/testing

	fmt.Println("Setting up server cache...")
	memoryCache := cache.NewLRUMemCache(5)

	fmt.Println("Registering server paths to server mux...")
	cardValidator := utils.NewCardValidator()
	handler := handlers.NewHandler(cardValidator, memoryCache)

	// Wrap handler with middleware
	http.HandleFunc("/validate", middleware.ValidateJSONHeaderRequest(handler.GetValidation))

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
