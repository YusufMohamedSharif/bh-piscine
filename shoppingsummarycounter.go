package piscine

import "strings"

func countItems(receipt string) map[string]int {
	// Initialize a map to store item counts
	itemCount := make(map[string]int)

	// Split the receipt into individual items
	items := strings.Fields(receipt)

	// Count the occurrences of each item
	for _, item := range items {
		// Increment the count of the item in the map
		itemCount[item]++
	}

	return itemCount
}
