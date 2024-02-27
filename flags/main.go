package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || containsHelpFlag(args) {
		printHelp()
		return
	}

	var insertString string
	var stringArg string
	var orderFlag bool

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "--insert", "-i":
			if i+1 < len(args) {
				insertString = extractInsertString(args[i+1])
				i++
			}
		case "--order", "-o":
			orderFlag = true
		default:
			stringArg = arg
		}
	}

	if insertString != "" {
		stringArg = insertStringIntoString(insertString, stringArg)
	}

	if orderFlag {
		stringArg = orderString(stringArg)
	}

	fmt.Println(stringArg)
}

func printHelp() {
	fmt.Println("Options:")
	fmt.Println("  --insert (-i) \t\tThis flag inserts the string in the next argument into the string given.")
	fmt.Println("\t\t\t\t  This flag must be followed by a string.")
	fmt.Println("  --order (-o)  \t\tThis flag orders the string argument.")
	fmt.Println("  --help  (-h)  \t\tThis flag prints the help message.")
}

func containsHelpFlag(args []string) bool {
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			return true
		}
	}
	return false
}

func extractInsertString(arg string) string {
	insertStr := ""
	for i := 0; i < len(arg); i++ {
		if arg[i] == '=' {
			insertStr = arg[i+1:]
			break
		}
	}
	return insertStr
}

func insertStringIntoString(insertString, mainString string) string {
	return mainString + insertString
}

func orderString(str string) string {
	if str == "" {
		return ""
	}
	minChar := 'z' + 1
	maxChar := 'a' - 1
	for _, char := range str {
		if char < minChar {
			minChar = char
		}
		if char > maxChar {
			maxChar = char
		}
	}
	result := ""
	for char := minChar; char <= maxChar; char++ {
		for _, c := range str {
			if c == char {
				result += string(c)
			}
		}
	}
	return result
}
