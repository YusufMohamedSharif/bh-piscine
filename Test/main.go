package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func PointOne(n *int) {
	*n = 1
}

func UltimatePointOne(n ***int) {
	***n = 1
}

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}

func UltimateDivMod(a *int, b *int) {
	temp := *a
	*a = *a / *b
	*b = temp % *b
}

func PrintStr(s string) {
	arr := []rune(s)

	for i := 0; i < len(arr); i++ {
		z01.PrintRune(arr[i])
	}
}

func StrLen(s string) int {
	arr := []rune(s)
	count := 0

	for i := 0; i <= len(arr); i++ {
		count += 1
	}
	z01.PrintRune(rune(count))
	return count
}

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func StrRev(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
