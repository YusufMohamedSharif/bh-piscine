package piscine

func AppendRange(min, max int) []int {
	result := []int{}
	direracne := max - min
	if min >= max {
		return result
	}
	for i := 0; i < direracne; i++ {
		result = append(result, i+direracne)
	}
	return result
}
