package models

type CardValidationResponse struct {
	IsValid bool   `json:"isValid"`
	Message string `json:"message,omitempty"`
}
