package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	for i := 1; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := i; k <= 9; k++ {
				for l := 0; l <= 9; l++ {
					if i != k || j != l {
						z01.PrintRune(rune(i) + '0')
						z01.PrintRune(rune(j) + '0')
						z01.PrintRune(' ')
						z01.PrintRune(rune(k) + '0')
						z01.PrintRune(rune(l) + '0')
						z01.PrintRune(',')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
