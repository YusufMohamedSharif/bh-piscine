package piscine

func IsUpper(s string) bool {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		if arr[i] > 'Z' || arr[i] < 'A' {
			return false
		}
	}
	return true
}
