package piscine

func ConcatParams(args []string) string {
	result := ""

	for i := 0; i < len(args); i++ {
		result += args[i] + string('\n')
	}
	return result
}
