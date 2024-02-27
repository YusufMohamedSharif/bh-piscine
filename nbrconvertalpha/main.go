package main

import (
	"os"
)

func main() {
	args := os.Args[1:]

	// Check if the --upper flag is present
	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		if n, valid := stringToInt(arg); valid && n >= 1 && n <= 26 {
			if upper {
				printChar('A' + byte(n-1))
			} else {
				printChar('a' + byte(n-1))
			}
		} else {
			printChar(' ')
		}
	}

	printChar('\n')
}

// stringToInt converts a string to an integer and returns true if successful, false otherwise
func stringToInt(s string) (int, bool) {
	num := 0
	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0, false
		}
		num = num*10 + int(digit-'0')
	}
	return num, true
}

// printChar prints a single character without using print or println functions
func printChar(c byte) {
	os.Stdout.Write([]byte{c})
}
