package piscine

func CapitalizeWords(s string) string {
	arr1 := []rune(s)
	arr2 := []rune{}
	newWordFlag := true // Flag to indicate the start of a new word
	for i := 0; i < len(arr1); i++ {
		if newWordFlag && ((arr1[i] >= 'a' && arr1[i] <= 'z') || (arr1[i] >= 'A' && arr1[i] <= 'Z') || (arr1[i] >= '0' && arr1[i] <= '9')) {
			if arr1[i] >= 'a' && arr1[i] <= 'z' {
				arr2 = append(arr2, arr1[i]-32) // Convert to uppercase
			} else {
				arr2 = append(arr2, arr1[i]) // Already uppercase, add as it is
			}
			newWordFlag = false
		} else {
			if arr1[i] >= 'A' && arr1[i] <= 'Z' {
				arr2 = append(arr2, arr1[i]+32) // Convert to lowercase
			} else {
				arr2 = append(arr2, arr1[i]) // Add the character as it is
			}
		}
		if !((arr1[i] >= 'a' && arr1[i] <= 'z') || (arr1[i] >= 'A' && arr1[i] <= 'Z') || (arr1[i] >= '0' && arr1[i] <= '9')) {
			newWordFlag = true // Next character should start a new word
		}
	}
	return string(arr2)
}
