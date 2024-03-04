package piscine

import "github.com/01-edu/z01"

func PrintNumber(a int) {
	z01.PrintRune(rune(a))
}

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a)-1; i++ {
		f(a[i])
	}
}
