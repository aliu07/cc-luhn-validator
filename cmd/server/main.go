package main

import (
	"cc-luhn-validator/internal/factory"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	serverLogger := log.New(os.Stdout, "[SERVER] ", log.LstdFlags)
	cacheLogger := log.New(os.Stdout, "[CACHE] ", log.LstdFlags)
	errorLogger := log.New(os.Stdout, "[ERROR] ", log.LstdFlags)

	serverLogger.Println("Setting up server instance...")
	config := factory.ServerConfig{
		ID:           "server1",
		CacheSize:    5,
		CacheTTL:     1 * time.Minute,
		ServerLogger: serverLogger,
		CacheLogger:  cacheLogger,
		ErrorLogger:  errorLogger,
	}
	server := factory.CreateServer(config)

	serverLogger.Println("Registering server paths to server mux...")
	server.SetupRoutes()

	// Specify IP address before colon to tell server to listen on specific IP addresses.
	serverLogger.Println("Listennig and ready to server...")
	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		serverLogger.Println("Server shutting down...")
	} else if err != nil {
		serverLogger.Fatalf("Error starting server: %s\n", err.Error())
	}
}
