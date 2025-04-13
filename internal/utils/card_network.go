package utils

import (
	"strconv"
)

type Network int

const (
	Visa Network = iota
	Mastercard
	AmericanExpress
	Unknown
)

func (n Network) String() string {
	switch n {
	case Visa:
		return "visa"
	case Mastercard:
		return "mastercard"
	case AmericanExpress:
		return "american express"
	default:
		return "unkown"
	}
}

// More information on card networks can be found here:
// https://en.wikipedia.org/wiki/Payment_card_number
func GetCardNetwork(cardNumber string) string {
	if len(cardNumber) < 4 {
		return Unknown.String()
	}

	prefix, err := strconv.Atoi(cardNumber[0:4])

	if err != nil {
		return Unknown.String()
	}

	switch {
	case prefix/1000 == 4:
		return Visa.String()
	case (prefix/100 >= 51 && prefix/100 <= 55) || (prefix >= 2221 && prefix <= 2720):
		return Mastercard.String()
	case prefix/100 == 34 || prefix/100 == 37:
		return AmericanExpress.String()
	default:
		return Unknown.String()
	}
}
