package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	// Initialize an array to store the count of each digit
	digitCount := [10]int{}

	// Count the occurrence of each digit
	for n > 0 {
		digit := n % 10
		digitCount[digit]++
		n /= 10
	}

	// Print digits in ascending order
	for i := 0; i < 10; i++ {
		for j := 0; j < digitCount[i]; j++ {
			z01.PrintRune(rune(i) + '0')
		}
	}
}
