package utils

import (
	"fmt"
	"strconv"
)

type Validator interface {
	ValidateDigits([]int) bool
	ValidateString(cardNumber string) (bool, error)
}

type CardValidator struct {
	// Configs
}

func NewCardValidator() *CardValidator {
	return &CardValidator{}
}

func (v *CardValidator) ValidateDigits(digits []int) bool {
	n := len(digits)

	if len(digits) == 0 {
		return false
	}

	sum := 0
	parity := n % 2

	for i, num := range digits[:n-1] {
		if i%2 != parity {
			sum += num
		} else if num > 4 {
			sum += 2*num - 9
		} else {
			sum += 2 * num
		}
	}

	return digits[n-1] == ((10 - (sum % 10)) % 10)
}

func (v *CardValidator) ValidateString(cardNumber string) (bool, error) {
	var cardNumbers []int

	for _, char := range cardNumber {
		digit, err := strconv.Atoi(string(char))

		if err != nil {
			return false, fmt.Errorf("Invalid digit in card number: %v", err)
		}

		cardNumbers = append(cardNumbers, digit)
	}

	return v.ValidateDigits(cardNumbers), nil
}
