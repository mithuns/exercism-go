package luhn

import "unicode"

func Valid(id string) bool {
	if len(id) <= 1 {
		return false
	}
	sum := 0
	newId := make([]rune, 0)
	for i := range id {
		r := rune(id[i])
		if unicode.IsDigit(r) {
			newId = append([]rune{r - '0'}, newId...)
		} else if !unicode.IsSpace(r) {
			return false
		}
	}
	if len(newId) <= 1 {
		return false
	}

	for i := 0; i < len(newId); i++ {
		if i%2 != 0 {
			val := int(newId[i])
			val = val * 2
			if val > 9 {
				val = val - 9
			}
			sum += val
		} else {
			sum += int(newId[i])
		}
	}
	return sum%10 == 0
}
