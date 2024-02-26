package piscine

func TrimAtoi(s string) int {
	arr1 := []rune(s)
	arr2 := []rune{}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] >= '0' && arr1[i] <= '9' || arr1[i] == '-' {
			arr2 = append(arr2, arr1[i])
		}
	}
	str := string(arr2)
	result := 0
	for _, ch := range str {
		if ch != '-' {
			digit := int(ch - '0')
			result = result*10 + digit
		}
	}
	if len(arr2) > 0 && arr2[0] == '-' {
		return -1 * result
	} else {
		return result
	}
}
