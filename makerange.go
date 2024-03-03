package piscine

func MakeRange(min, max int) []int {
	diferance := max - min
	if min >= max {
		return nil
	}
	result := make([]int, diferance)
	for i := min; i < max; i++ {
		result[i-min] = i
	}
	return result
}
