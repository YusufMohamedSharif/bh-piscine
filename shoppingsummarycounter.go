package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	// Initialize a map to store item counts
	itemCount := make(map[string]int)

	// Initialize variables for tracking item and word boundaries
	itemStart := 0
	inWord := false

	// Iterate through the receipt string
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' && !inWord {
			// Start of a new item
			inWord = true
			itemStart = i
		} else if str[i] == ' ' && inWord {
			// End of the current item
			inWord = false
			// Extract the item substring and increment its count in the map
			item := str[itemStart:i]
			itemCount[item]++
		}
	}

	// Handle the last item in the receipt string
	if inWord {
		item := receipt[itemStart:]
		itemCount[item]++
	}

	return itemCount
}
