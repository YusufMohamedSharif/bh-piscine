package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 {
		return 0 // Empty string, return 0
	}

	sign := 1 // Sign of the integer, default is positive
	result := 0
	foundNumber := false

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			foundNumber = true
			digit := int(ch - '0')
			result = result*10 + digit
		} else if ch == '-' && !foundNumber {
			sign = -1 // Found a negative sign before any number
		} else if ch != ' ' {
			return 0 // Invalid input, return 0
		}
	}

	return result * sign
}
