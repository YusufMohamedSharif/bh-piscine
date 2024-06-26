package piscine

func Unmatch(numbers []int) int {
	counts := make(map[int]int)

	// Count occurrences of each number
	for _, num := range numbers {
		counts[num]++
	}

	// Check for the unpaired number
	for num, count := range counts {
		if count%2 != 0 {
			return num // Return the unpaired number
		}
	}

	// If no unpaired number is found, return -1
	return -1
}
