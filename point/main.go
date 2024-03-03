package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func P(x rune) {
	z01.PrintRune(x)
}

func main() {
	points := &point{}

	setPoint(points)

	xS := "x = "
	yS := ", y = "

	for _, a := range xS {
		z01.PrintRune(a)
	}
	P(52)
	P(50)

	for _, b := range yS {
		z01.PrintRune(b)
	}
	P(50)
	P(49)
	P('\n')
}
