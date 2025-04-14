package models

type CardValidationResponse struct {
	IsValid     bool   `json:"isValid"`
	CardNetwork string `json:"cardNetwork"`
	Message     string `json:"message,omitempty"`
	Source      string `json:"source"`
}
