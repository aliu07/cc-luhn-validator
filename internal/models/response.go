package models

type CardValidationResponse struct {
	IsValid     bool   `json:"isValid"`
	CardNetwork string `json:"cardNetwork"`
	Source      string `json:"source"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
