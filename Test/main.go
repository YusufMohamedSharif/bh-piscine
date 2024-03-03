package main

import (
	"fmt"

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

// //////Quest 4, 5, 6//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func IterativeFactoriall(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else {
		result := 1
		for i := 1; i <= nb; i++ {
			result *= i
		}
		return result
	}
}

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else if nb <= 1 {
		return 1
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}

func IterativePower(nb int, power int) int {
	result := nb
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		for i := 1; i < power; i++ {
			result *= nb
		}
	}
	return result
}

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else if power == 1 {
		return nb
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	root := 0
	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			root = i
			break
		}
	}
	return root
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

func FirstRune(s string) rune {
	arr := []rune(s)
	return arr[0]
}

func LastRune(s string) rune {
	arr := []rune(s)
	return arr[len(arr)-1]
}

func Concat(str1 string, str2 string) string {
	return str1 + str2
}

func main() {
	// find the index of the last '/'

	// fmt.Println(piscine.AppendRange(0, 1))
	fmt.Println(piscine.MakeRange(0, 1))

	// fmt.Println(piscine.Join([]string{"Hello!", " How", " are", " you?"}, "123"))
}
