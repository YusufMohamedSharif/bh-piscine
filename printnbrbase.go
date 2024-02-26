package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(n int, base string) {
	isValidBase := func(base string) bool {
		if len(base) < 2 {
			return false
		}

		seen := make(map[rune]bool)
		for _, char := range base {
			if char == '+' || char == '-' {
				return false
			}
			if seen[char] {
				return false
			}
			seen[char] = true
		}

		return true
	}

	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		// z01.Flush()
		return
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n *= -1
	}

	var result []rune
	for n > 0 {
		remainder := n % len(base)
		result = append(result, rune(base[remainder]))
		n /= len(base)
	}

	if isNegative {
		z01.PrintRune('-')
	}

	// Reverse the result slice
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	for _, char := range result {
		z01.PrintRune(char)
	}

	// z01.Flush()
}
