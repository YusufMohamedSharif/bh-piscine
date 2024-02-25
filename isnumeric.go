package piscine

func IsNumeric(s string) bool {
	arr := []rune(s)
	if len(arr) == 0 {
		return true
	}
	for i := 0; i < len(arr); i++ {
		if arr[i] >= '0' && arr[i] <= '9' {
		} else {
			return false
		}
	}
	return true
}
