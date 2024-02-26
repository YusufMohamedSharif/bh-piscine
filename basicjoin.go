package piscine

func BasicJoin(strs []string) string {
	result := ""
	for i := 0; i < len(strs); i++ {
		result += strs[i]
	}
	return result
}
