package piscine

import "math"

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(nb)))
	for i := 5; i <= sqrtN; i += 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}
