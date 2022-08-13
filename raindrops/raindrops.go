package raindrops

import "strconv"

func Convert(number int) string {
	outcome := ""
	if number%3 == 0 {
		outcome += "Pling"
	}
	if number%5 == 0 {
		outcome += "Plang"
	}
	if number%7 == 0 {
		outcome += "Plong"
	}
	if outcome == "" {
		outcome = strconv.Itoa(number)
	}
	return outcome
}
