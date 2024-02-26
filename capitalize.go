package piscine

func Capitalize(s string) string {
	arr1 := []rune(s)
	arr2 := []rune{}
	newWordFlag := false
	for i := 0; i < len(arr1); i++ {
		if (i == 0 || newWordFlag) && (arr1[i] >= 'a' && arr1[i] <= 'z') {
			arr2 = append(arr2, arr1[i]-32)
			newWordFlag = false
		} else {
			arr2 = append(arr2, arr1[i])
		}
		if arr1[i] < 'A' || (arr1[i] > 'Z' && arr1[i] < 'a') || arr1[i] > 'z' {
			newWordFlag = true
		}
	}
	return string(arr2)
}
