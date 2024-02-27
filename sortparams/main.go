package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Get the command-line arguments excluding the program name
	args := os.Args[1:]

	// Sort the arguments in ASCII order using bubble sort
	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-i-1; j++ {
			if compare(args[j], args[j+1]) > 0 {
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Print the sorted arguments
	for _, arg := range args {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

// compare returns a negative value if a is less than b, 0 if they are equal, and a positive value if a is greater than b
func compare(a, b string) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}
	return 0
}
