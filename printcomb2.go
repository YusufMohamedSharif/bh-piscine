package piscine

import (
	//"github.com/01-edu/z01"
	"fmt"
)

func PrintComb2() {
	for i := 10; i < 100; i++ {
		for j := i + 1; j < 100; j++ {
			fmt.Printf("%02d %02d ", i, j)
		}
	}
}
