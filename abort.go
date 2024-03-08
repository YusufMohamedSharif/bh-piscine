package piscine

import "sort"

func Abort(a, b, c, d, e int) int {
	// Store the numbers in a slice
	numbers := []int{a, b, c, d, e}

	// Sort the numbers in ascending order
	sort.Ints(numbers)

	// Check if the number of elements is odd or even
	n := len(numbers)
	if n%2 == 1 {
		// If odd, return the middle number
		return numbers[n/2]
	} else {
		// If even, return the average of the two middle numbers
		middleLeft := numbers[n/2-1]
		middleRight := numbers[n/2]
		return (middleLeft + middleRight) / 2.0
	}
}
