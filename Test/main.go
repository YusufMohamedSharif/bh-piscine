package main

import (
	piscine "bh-piscine"

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

/*func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '8'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					z01.PrintRune(i)
					z01.PrintRune(j)
					z01.PrintRune(' ')
					z01.PrintRune(k)
					if i == '0' && j == '0' && k == '0' && l == '0' {
						l++
						z01.PrintRune(l)
					} else {
						z01.PrintRune(l)
					}

					if i == '9' && j == '8' && k == '9' && l == '9' {
						break
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}*/

func PrintComb2() {
	for i := '1'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := i; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					if i != k || j != l {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 1 || nb > 20 {
		return 0
	} else {

		for i := 1; i <= nb; i++ {
			result = result * i
		}
		return result
	}
}

func main() {
	piscine.PrintNbrInOrder(56423147)
	// piscine.PrintNbr(-9223372036854775808)
}
