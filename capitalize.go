package piscine

func Capitalize(s string) string {
	arr1 := []rune(s)
	arr2 := []rune{}
	capitalizeNext := true // Flag to indicate whether the next character should be capitalized
	for i := 0; i < len(arr1); i++ {
		if capitalizeNext && ((arr1[i] >= 'a' && arr1[i] <= 'z') || (arr1[i] >= 'A' && arr1[i] <= 'Z')) {
			if arr1[i] >= 'a' && arr1[i] <= 'z' {
				arr2 = append(arr2, arr1[i]-32) // Convert to uppercase
			} else {
				arr2 = append(arr2, arr1[i]) // Already uppercase, add as it is
			}
			capitalizeNext = false
		} else {
			if arr1[i] >= 'A' && arr1[i] <= 'Z' {
				arr2 = append(arr2, arr1[i]+32) // Convert to lowercase
			} else {
				arr2 = append(arr2, arr1[i]) // Add the character as it is
			}
		}
		if !((arr1[i] >= 'a' && arr1[i] <= 'z') || (arr1[i] >= 'A' && arr1[i] <= 'Z')) {
			capitalizeNext = true // Next character should be capitalized if it's alphanumeric
		}
	}
	return string(arr2)
}