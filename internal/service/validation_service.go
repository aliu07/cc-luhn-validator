package service

import (
	"cc-luhn-validator/internal/cache"
	"cc-luhn-validator/internal/constants"
	"cc-luhn-validator/internal/utils"
	"errors"
	"time"
)

type CardValidationService interface {
	ValidateCard(cardNumber string) (ValidationResult, error)
}

type cardValidationService struct {
	validator utils.Validator
	cache     cache.Cache
	cacheTTL  time.Duration
}

type ValidationResult struct {
	IsValid     bool
	CardNetwork string
	Source      string
}

func NewCardValidationService(v utils.Validator, c cache.Cache, ttl time.Duration) CardValidationService {
	return &cardValidationService{
		validator: v,
		cache:     c,
		cacheTTL:  ttl,
	}
}

func (s *cardValidationService) ValidateCard(cardNumber string) (ValidationResult, error) {
	if cardNumber == "" {
		return ValidationResult{}, errors.New("Missing card number")
	}

	// Check cache
	if data, exists := s.cache.Get(cardNumber); exists {
		return ValidationResult{
			IsValid:     data.IsValid,
			CardNetwork: data.CardNetwork,
			Source:      constants.Cache.String(),
		}, nil
	}

	isValid, err := s.validator.ValidateString(cardNumber)

	if err != nil {
		return ValidationResult{}, err
	}

	cardNetwork := utils.GetCardNetwork(cardNumber)

	// Put in cache
	s.cache.Put(cardNumber, isValid, cardNetwork, s.cacheTTL)

	return ValidationResult{
		IsValid:     isValid,
		CardNetwork: cardNetwork,
		Source:      constants.Server.String(),
	}, nil
}
