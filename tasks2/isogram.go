package tasks2

import (
	"strings"
)

func IsIsogram(word string) bool {
	word = strings.ToLower(word)

	for i := 0; i < len(word); i++ {
		for j := i + 1; j < len(word); j++ {
			if word[i] == word[j] && (word[i] != '-' && word[i] != ' ') {
				return false
			}
		}
	}
	return true

}

//	OR

// func IsIsogram(s string) bool {
// 	s = strings.ToLower(s)
// 	for i, c := range s {
// 		if unicode.IsLetter(c) && strings.ContainsRune(s[i+1:], c) {
// 			return false
// 		}
// 	}
// 	return true
// }