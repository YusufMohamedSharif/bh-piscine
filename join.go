package piscine

func Join(strs []string, sep string) string {
	result := ""
	for i := 0; i < len(strs); i++ {
		result += strs[i]
		result += sep
	}
	return result
}
