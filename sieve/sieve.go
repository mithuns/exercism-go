package sieve

func Sieve(limit int) []int {
	candidateDigitsUptoLimit := make([]bool, limit+1)
	for i := 2; i <= limit; i++ {
		candidateDigitsUptoLimit[i] = true
	}
	for i := 2; i*i <= limit; i++ {
		if candidateDigitsUptoLimit[i] {
			multiplicationFactor := 0
			for j := i * i; j <= limit; j = (i * i) + (i * multiplicationFactor) {
				candidateDigitsUptoLimit[j] = false
				multiplicationFactor++
			}
		}
	}
	output := make([]int, 0)
	for number := range candidateDigitsUptoLimit {
		if candidateDigitsUptoLimit[number] {
			output = append(output, number)
		}
	}
	return output
}
