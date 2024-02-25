package piscine

func AlphaCount(s string) int {
	arr := []rune(s)
	counter := 0
	for i := 0; i <= len(arr)-1; i++ {
		if (arr[i] >= 'a' && arr[i] <= 'z') || (arr[i] >= 'A' && arr[i] <= 'Z') {
			counter++
		}
	}
	return counter
}
