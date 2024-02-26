package piscine

func Capitalize(s string) string {
	arr1 := []rune(s)
	arr2 := []rune{}
	newWordFlag := false
	for i := 0; i < len(arr1); i++ {
		if i == 0 || newWordFlag && (arr1[i] >= 'a' && arr1[i] <= 'z') {
			arr2 = append(arr2, arr1[i]-32)
		} else {
			arr2 = append(arr2, arr1[i])
		}
		newWordFlag = false
		if arr1[i] == ' ' || arr1[i] == '+' {
			newWordFlag = true
		}
	}
	return string(arr2)
}
