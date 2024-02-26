package piscine

func Capitalize(s string) string {
	arr1 := []rune(s)
	arr2 := []rune{}
	shouldCapitalize := true
	for i := 0; i < len(arr1); i++ {
		if shouldCapitalize && (arr1[i] >= 'a' && arr1[i] <= 'z') {
			arr2 = append(arr2, arr1[i]-32) // Convert to uppercase
			shouldCapitalize = false
		} else if !shouldCapitalize && (arr1[i] >= 'A' && arr1[i] <= 'Z') {
			arr2 = append(arr2, arr1[i]+32) // Convert to lowercase
		} else {
			arr2 = append(arr2, arr1[i]) // Add the character as it is
		}
		if !((arr1[i] >= 'a' && arr1[i] <= 'z') || (arr1[i] >= 'A' && arr1[i] <= 'Z')) {
			shouldCapitalize = true
		}
	}
	return string(arr2)
}
