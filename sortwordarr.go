package piscine

func SortWordArr(arr []string) {
	// Implement a simple sorting algorithm like Bubble Sort or Insertion Sort
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if compareStrings(arr[j], arr[j+1]) > 0 {
				// Swap the elements if they are in the wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// compareStrings compares two strings by their ASCII values
func compareStrings(a, b string) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return int(a[i]) - int(b[i])
		}
	}
	return len(a) - len(b)
}
