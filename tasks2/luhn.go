package tasks2

import (
	"strconv"
	"unicode"
)

func Valid(id string) bool {

	cleaned := ""
	for _, r := range id {
		if unicode.IsDigit(r) {
			cleaned += string(r)
		} else if r != ' ' {
			return false 
		}
	}

	if len(cleaned) <= 1 {
		return false
	}

	sum := 0
	double := false


	for i := len(cleaned) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(cleaned[i]))
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}
