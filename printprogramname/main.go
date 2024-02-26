package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	temp := ""
	result := ""
	programName := os.Args[:1]
	args := programName[0]
	for i := len(args) - 1; i >= 0; i-- {
		if args[i] == '/' {
			break
		} else {
			temp += string(args[i])
		}
	}
	for i := len(temp) - 1; i >= 0; i-- {
		result += string(temp[i])
	}
	for i := 0; i < len(result); i++ {
		z01.PrintRune(rune(result[i]))
	}

	z01.PrintRune('\n')
}
