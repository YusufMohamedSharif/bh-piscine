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

func main() {
	points := &point{}

	setPoint(points)

	xS := "x = "
	yS := ", y = "

	for _, a := range xS {
		z01.PrintRune(a)
	}
	z01.PrintRune(52)
	z01.PrintRune(50)

	for _, b := range yS {
		z01.PrintRune(b)
	}
	z01.PrintRune(50)
	z01.PrintRune(49)
	z01.PrintRune('\n')
}
