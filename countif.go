package piscine

func CountIf(f func(string) bool, tab []string) int {
	var result bool
	var counter int
	for _, number := range tab {
		result = f(number)
		if result {
			counter++
		}
	}
	return counter
}
