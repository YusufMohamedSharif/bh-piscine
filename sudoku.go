package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:] // Ignore the filename
	if checkInput(arguments) {
		table := [9][9]rune{}
		table = fillTable(table, arguments)
		if isSolve(&table, 0) {
			for y := 0; y < 9; y++ {
				for x := 0; x < 9; x++ {
					z01.PrintRune(rune(table[y][x]))
					if x != 8 {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}
		} else {
			fmt.Println("Error") // Multiple solutions or no solution
		}
	}
}

func checkInput(args []string) bool {
	if len(args) != 9 {
		fmt.Println("Error") // Input length is out of range, wrong number of rows
		return false
	}
	for i := 0; i < len(args); i++ {
		if len(args[i]) != 9 {
			fmt.Println("Error") //  Input length is out of range , wrong number of columns
			return false
		}
	}
	for i := 0; i < len(args); i++ {
		for _, value := range args[i] {
			if (value >= '1' && value <= '9') || value == '.' {
			} else {
				fmt.Println("Error") // Error, invalid input
				return false
			}
		}
	}
	return true
}

// fills with input arguments
func fillTable(table [9][9]rune, args []string) [9][9]rune {
	for i := range args {
		for j := range args[i] {
			table[i][j] = rune(args[i][j])
		}
	}
	return table
}

// check if it fits
func isValid(table *[9][9]rune, x int, y int, z rune) bool {
	// check double int
	for i := 0; i < 9; i++ {
		if z == table[i][x] {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if z == table[y][j] {
			return false
		}
	}
	// square check
	a := x / 3
	b := y / 3
	for k := 3 * a; k < 3*(a+1); k++ {
		for l := 3 * b; l < 3*(b+1); l++ {
			if z == table[l][k] {
				return false
			}
		}
	}
	return true
}

// backtracking solver
func isSolve(table *[9][9]rune, count int) bool {
	if count > 1 {
		return false // Multiple solutions found
	}
	if !isDots(table) {
		return true
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if table[y][x] == '.' {
				for z := '1'; z <= '9'; z++ {
					if isValid(table, x, y, z) {
						table[y][x] = z
						if isSolve(table, count+1) {
							return true
						}
					}
					table[y][x] = '.'
				}
				return false
			}
		}
	}
	return false
}

// counts remaining empty cells
func isDots(table *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == '.' {
				return true
			}
		}
	}
	return false
}
