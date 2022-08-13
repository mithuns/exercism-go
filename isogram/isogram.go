package isogram

import (
	"strings"
	"unicode"
)

/*func IsIsogram(word string) bool {
	if len(word) == 0 {
		return true
	}
	word = strings.ToLower(word)
	runeMap := make(map[rune]int, 0)
	for _, c := range word {
		if string(c) != " " && string(c) != "-" {
			if _, found := runeMap[c]; found {
				return false
			} else {
				runeMap[c] = 1
			}
		}
	}
	return true
}*/
func IsIsogram(word string) bool {
	s := strings.ToLower(word)
	for _, r := range s {
		if unicode.IsLetter(r) && strings.Count(s, string(r)) > 1 {
			return false
		}
	}
	return true
}
