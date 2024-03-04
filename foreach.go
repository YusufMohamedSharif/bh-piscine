package piscine

import "fmt"

func PrintNumber(a int) {
	fmt.Print(a)
}

func ForEach(f func(int), a []int) {
	for i := 0; i < len(a)-1; i++ {
		f(a[i])
	}
}
