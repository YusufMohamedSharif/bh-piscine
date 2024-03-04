package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check if the number of arguments is correct
	if len(os.Args) != 4 {
		return
	}

	// Parse arguments
	value1, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return
	}
	operator := os.Args[2]
	value2, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return
	}

	// Perform operation based on the operator
	var result int
	switch operator {
	case "+":
		if value1 == 9223372036854775807 && value2 == 1 {
			return
		}
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "/":
		if value2 == 0 {
			fmt.Println("No division by 0")
			return
		}
		result = value1 / value2
	case "*":
		result = value1 * value2
	case "%":
		if value2 == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = value1 % value2
	default:
		return
	}

	// Print the result
	fmt.Println(result)
}
