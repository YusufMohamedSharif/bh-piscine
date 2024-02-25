package piscine

func Concat(str1 string, str2 string) string {
	arr := []rune(str2)
	for i := 0; i < len(str2); i++ {
		str1 += string(arr[i])
	}
	return str1
}
