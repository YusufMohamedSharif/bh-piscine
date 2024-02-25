package piscine

/*func NRune(s string, n int) rune {
	if n < 0 || n >= len(s) {
		return 0
	}
	return []rune(s)[n]
}*/

func NRune(s string, n int) rune {
	if n >= 0 && n < len(s) {
		return rune(s[n])
	} else {
		return 0
	}
}
