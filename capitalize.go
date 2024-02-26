package piscine

func Capitalize(s string) string {
	arr1 := []rune(s)
	arr2 := []rune{}
	newWordFlag := true // Set to true to capitalize the first letter of the string
	for i := 0; i < len(arr1); i++ {
		if newWordFlag && (arr1[i] >= 'a' && arr1[i] <= 'z') {
			arr2 = append(arr2, arr1[i]-32) // Convert to uppercase
			newWordFlag = false
		} else if newWordFlag && (arr1[i] >= 'A' && arr1[i] <= 'Z') {
			arr2 = append(arr2, arr1[i]) // Already uppercase, add as it is
			newWordFlag = false
		} else if (arr1[i] < 'a' || arr1[i] > 'z') && (arr1[i] < 'A' || arr1[i] > 'Z') {
			newWordFlag = true // Found a non-letter character, next character is the start of a new word
			arr2 = append(arr2, arr1[i]) // Add the non-letter character as it is
		} else {
			arr2 = append(arr2, arr1[i]) // Add the character as it is
		}
	}
	return string(arr2)
}