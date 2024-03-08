package piscine

func Abort(a, b, c, d, e int) int {
	// Bubble sort the numbers
	numbers := []int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := 0; j < 4-i; j++ {
			if numbers[j] > numbers[j+1] {
				// Swap the numbers if they're in the wrong order
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	// Calculate the median
	middleIndex := 5 / 2
	if 5%2 == 1 {
		// If odd, return the middle number
		return numbers[middleIndex]
	} else {
		// If even, return the average of the two middle numbers
		middleLeft := numbers[middleIndex-1]
		middleRight := numbers[middleIndex]
		return (middleLeft + middleRight) / 2.0
	}
}
