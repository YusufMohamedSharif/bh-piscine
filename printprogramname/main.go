package main

import (
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]

	// Find the index of the last '/' in the program path
	lastSlashIndex := strings.LastIndex(programName, "/")

	// Extract the program name from the path
	if lastSlashIndex != -1 {
		programName = programName[lastSlashIndex+1:]
	}

	// Print the program name character by character
	for _, char := range programName {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}

/*func main() {

	temp := ""
	result := ""
	// programName := os.Args[:1]
	// args := programName[0]
	programName := os.Args[0]

	for i := len(programName) - 1; i >= 0; i-- {
		if programName[i] == '/' {
			break
		} else {
			temp += string(programName[i])
		}
	}
	for i := len(temp) - 1; i >= 0; i-- {
		result += string(temp[i])
	}
	for i := 0; i < len(result); i++ {
		z01.PrintRune(rune(result[i]))
	}

	z01.PrintRune('\n')
}*/

/*
func main() {
	programName := os.Args[0]

	for i := 0; i < len(programName); i++ {
		if programName[i] != '/' && programName[i] != '.' {
			z01.PrintRune(rune(programName[i]))
		}
	}
}*/
