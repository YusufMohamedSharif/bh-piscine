package main

import (
	"fmt"
)

func PointOne(n *int) {
	*n = 1
}

func UltimatePointOne(n ***int) {
	***n = 1
}

func main() {
	n := 0
	a := &n
	b := &a
	UltimatePointOne(&b)

	fmt.Println(n)
}
