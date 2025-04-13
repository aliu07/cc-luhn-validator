package main

import (
	"cc-luhn-validator/internal"
	"cc-luhn-validator/internal/validation"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Dummy credit card numbers can be found in stripe testing doc: https://docs.stripe.com/testing

	fmt.Println("Starting server...")

	fmt.Println("Registering server paths to server mux...")
	cardValidator := validation.NewCardValidator()
	handler := internal.NewHandler(cardValidator)
	http.HandleFunc("/validate", handler.GetValidation)

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
