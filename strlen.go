package piscine

func StrLen(s string) int {
	arr := []rune(s)
	count := 0

	for i := 0; i < len(arr); i++ {
		count += 1
	}
	return count
}