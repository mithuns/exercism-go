package prime

func Factors(n int64) []int64 {
	factors := make([]int64, 0)
	if n%2 == 0 {
		factors = append(factors, 2)
	}
	for n%2 == 0 {
		n = n / 2
	}
	for i := int64(3); i*i <= n; i = i + 2 {
		if n%i == 0 {
			factors = append(factors, i)
		}
		for n%i == 0 {
			n = n / i
		}
	}
	if n > 2 {
		factors = append(factors, n)
	}
	return factors

}
