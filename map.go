package piscine

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, number := range a {
		result = append(result, f(number))
	}
	return result
}
