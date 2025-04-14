package constants

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
