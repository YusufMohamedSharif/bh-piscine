package piscine

func BasicJoin(strs ...string) string {
	totalLength := 0
	for _, str := range strs {
		totalLength += len(str)
	}

	result := make([]byte, totalLength)
	index := 0
	for _, str := range strs {
		for i := 0; i < len(str); i++ {
			result[index] = str[i]
			index++
		}
	}

	return string(result)
}
