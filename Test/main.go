package main

import (
	"fmt"
)

func PointOne(n *int) {
	*n = 1
}

func UltimatePointer(n ***int) {
	***n = 1
}

func main() {
	n := 0
	a := &n
	b := &a
	UltimatePointer(&b)

	fmt.Println(n)
}
