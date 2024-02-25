package piscine

func IsAlpha(s string) bool {
	arr := []rune(s)
	if len(arr) == 0 {
		return true
	}
	for i := 0; i < len(arr); i++ {
		if (arr[i] >= 'a' && arr[i] <= 'z') || (arr[i] >= 'A' && arr[i] <= 'Z') || (arr[i] >= '0' && arr[i] <= '9') {
		} else {
			return false
		}
	}
	return true
}
