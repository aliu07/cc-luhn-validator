package src

// The following function takes a card number, including the check digit, as an array
// of integers and outputs true if the check digit is correct, false otherwise.
func IsValid(cardNumbers []int) bool {
	n := len(cardNumbers)
	sum := 0
	parity := n % 2

	for i, num := range cardNumbers[:n-1] {
		if i%2 != parity {
			sum += num
		} else if num > 4 {
			sum += 2*num - 9
		} else {
			sum += 2 * num
		}
	}

	return cardNumbers[n-1] == ((10 - (sum % 10)) % 10)
}
