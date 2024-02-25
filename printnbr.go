package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int64) {
	if n == -9223372036854775808 {
		// Handle special case of minimum int value separately
		z01.PrintRune('-')
		z01.PrintRune('9')
		n = 223372036854775808 // Set n to positive value with overflow adjusted
	} else if n < 0 {
		z01.PrintRune('-')
		n = n * -1
	} else if n == 0 {
		z01.PrintRune('0')
	}
	digits := ""
	for i := n; i > 0; i = i / 10 {
		digits += string((rune((i % 10) + '0')))
	}

	for j := len(digits) - 1; j >= 0; j-- {
		z01.PrintRune(rune(digits[j]))
	}

	z01.PrintRune('\n')
}
