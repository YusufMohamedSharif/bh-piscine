package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	} else if power == 1 || power == 0 {
		return nb
	} else if power == 0 {
		return 1
	} else {
		for i := 0; i < power; i++ {
			result *= nb
		}
	}
	return result
}
