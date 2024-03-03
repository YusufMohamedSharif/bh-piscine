package piscine

func SplitWhiteSpaces(s string) []string {
	arr := []string{}
	result := ""
	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if result != "" {
				arr = append(arr, result)
			}
			result = ""
		} else {
			result += string(char)
		}
	}
	if result != "" {
		arr = append(arr, result)
	}

	return arr
}
