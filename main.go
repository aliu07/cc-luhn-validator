package main

import (
	"cc-luhn-validator/src"
	"fmt"
)

func main() {
	// Numbers taken from stripe testing doc: https://docs.stripe.com/testing

	// Invalid credit card
	invalidCardNumbers := []int{4, 2, 4, 2, 4, 2, 4, 2, 4, 2, 4, 2, 4, 2, 4, 1}
	isValid := src.IsValid(invalidCardNumbers)
	fmt.Printf("Card is valid: %v\n", isValid)

	// Valid credit card
	validCardNumbers := []int{4, 2, 4, 2, 4, 2, 4, 2, 4, 2, 4, 2, 4, 2, 4, 2}
	isValid = src.IsValid(validCardNumbers)
	fmt.Printf("Card is valid: %v\n", isValid)
}
