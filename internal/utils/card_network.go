package utils

import (
	"cc-luhn-validator/internal/constants"
	"strconv"
)

// More information on card networks can be found here:
// https://en.wikipedia.org/wiki/Payment_card_number
func GetCardNetwork(cardNumber string) string {
	if len(cardNumber) < 4 {
		return constants.Unknown.String()
	}

	prefix, err := strconv.Atoi(cardNumber[0:4])

	if err != nil {
		return constants.Unknown.String()
	}

	switch {
	case prefix/1000 == 4:
		return constants.Visa.String()
	case (prefix/100 >= 51 && prefix/100 <= 55) || (prefix >= 2221 && prefix <= 2720):
		return constants.Mastercard.String()
	case prefix/100 == 34 || prefix/100 == 37:
		return constants.AmericanExpress.String()
	default:
		return constants.Unknown.String()
	}
}
