package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)
	if n <= len(arr) && n > 0 {
		return arr[n-1]
	} else {
		return 0
	}
}

/*func NRune(s string, n int) rune {
	if n < 0 || n >= len(s) {
		return 0
	}
	return []rune(s)[n]
}

func NRune(s string, n int) rune {
	if n > 0 && n <= len(s) {
		return rune(s[n-1])
	} else {
		return 0
	}
}*/
