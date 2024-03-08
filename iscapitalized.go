package piscine

func IsCapitalized(s string) bool {
	sep := true
	for _, ch := range s {
		if (ch >= 'a' && ch <= 'z') && sep {
			return false
		} else {
			sep = false
		}
		if !((ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')) {
			sep = true
		}
	}
	return true
}
