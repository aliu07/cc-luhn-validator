package src

import (
	"io"
	"net/http"
	"strconv"
)

func GetValidation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		cardNumberStr := r.URL.Query().Get("card")

		if cardNumberStr == "" {
			http.Error(w, "Missing card number", http.StatusBadRequest)
			return
		}

		var cardNumbers []int

		for _, char := range cardNumberStr {
			digit, err := strconv.Atoi(string(char))

			if err != nil {
				http.Error(w, "Invalid card number format", http.StatusBadRequest)
				return
			}

			cardNumbers = append(cardNumbers, digit)
		}

		isValid := IsValid(cardNumbers)
		io.WriteString(w, strconv.FormatBool(isValid)+"\n")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
