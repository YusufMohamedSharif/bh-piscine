package piscine

func Any(f func(string) bool, a []string) bool {
	var result bool
	for _, number := range a {
		result = f(number)
		if result {
			return true
		}
	}
	return false
}
