package piscine

import "github.com/01-edu/z01"

func printCombinations(n int, current []rune, start rune) {
	if n == 0 {
		for i, digit := range current {
			if i > 0 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			z01.PrintRune(digit)
		}
		z01.PrintRune('\n')
		return
	}

	for i := start; i <= '9'; i++ {
		newCurrent := append(current, i)
		printCombinations(n-1, newCurrent, i+1)
	}
}
