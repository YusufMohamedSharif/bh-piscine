package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)
	if n <= len(arr) && n > 0 {
		return arr[n-1]
	} else {
		return 0
	}
}
