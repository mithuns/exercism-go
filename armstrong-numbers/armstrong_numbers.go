package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	if n == 0 {
		return true
	}
	stringRepresentation := strconv.Itoa(n)
	exponent := len(stringRepresentation)
	result := 0
	for i := 0; i < len(stringRepresentation); i++ {
		if stringRepresentation[i] == '0' {
			result += 1
		} else {
			result += int(math.Pow(float64(stringRepresentation[i]-'0'), float64(exponent)))
		}
	}
	return result == n
}
