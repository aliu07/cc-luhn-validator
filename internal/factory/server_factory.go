package factory

import (
	"cc-luhn-validator/internal/cache"
	"cc-luhn-validator/internal/handlers"
	"cc-luhn-validator/internal/server"
	"cc-luhn-validator/internal/service"
	"cc-luhn-validator/internal/utils"
	"log"
	"time"
)

type ServerConfig struct {
	ID           string
	CacheSize    int
	CacheTTL     time.Duration
	ServerLogger *log.Logger
	CacheLogger  *log.Logger
	ErrorLogger  *log.Logger
}

func CreateServer(config ServerConfig) server.CardValidationServer {
	cardValidator := utils.NewCardValidator()

	memoryCache := cache.NewLRUMemCache(config.CacheSize, config.CacheLogger)
	cacheTTL := config.CacheTTL

	service := service.NewCardValidationService(
		cardValidator,
		memoryCache,
		cacheTTL,
	)

	handler := handlers.NewHandler(service, config.ErrorLogger)

	return server.NewValidationServer(config.ID, handler, config.ServerLogger, config.CacheLogger, config.ErrorLogger)
}
