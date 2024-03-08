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

// ////Quest 9, 8, 7////////////////////////////////////////////////////////////
func ForEach(f func(int), a []int) {
	for _, number := range a {
		f(number)
	}
}

func Map(f func(int) bool, a []int) []bool {
	var result []bool
	for _, number := range a {
		result = append(result, f(number))
	}
	return result
}

func Any(f func(string) bool, a []string) bool {
	for _, number := range a {
		if f(number) {
			return true
		}
	}
	return false
}

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, number := range tab {
		if f(number) {
			count++
		}
	}
	return count
}

func CompareInt(a, b int) int {
	if a == b {
		return 0
	} else if a < b {
		return -1
	} else {
		return 1
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func F(a int) {
	a = a + 1
	fmt.Println(a)
}

func ForEachh(f func(int), a []int) {
	for i := 0; i < len(a); i++ {
		f(a[i])
	}
}

func Anyy(f func(string) bool, a []string) bool {
	for i := 0; i < len(a); i++ {
		if f(a[i]) == true {
			return true
		}
	}
	return false
}

func CountIff(f func(string) bool, tab []string) int {
	counter := 0
	for i := 0; i < len(tab); i++ {
		if f(tab[i]) == true {
			counter++
		}
	}
	return counter
}

func Mapp(f func(int) bool, a []int) []bool {
	var result []bool
	for i := 0; i < len(a); i++ {
		result = append(result, f(a[i]))
	}
	return result
}

func AppendRange(min, max int) []int {
	var result []int
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	result := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		result[i] = i + min
	}
	return result
}

func ConcatParams(args []string) string {
	result := ""

	for i := 0; i < len(args); i++ {
		result += args[i]
		if i != len(args)-1 {
			result += "\n"
		}

	}
	return result
}

func SplitWhiteSpaces(s string) []string {
	var arr []string
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '\n' {
			if result != "" {
				arr = append(arr, result)
			}
			result = ""
		} else {
			result += string(s[i])
		}
	}
	if result != "" {
		arr = append(arr, result)
	}

	return arr
}

func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, char := range word {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func Atoi(s string) int {
	out := 0
	sign := 1

	for i, char := range s {
		if char >= '0' && char <= '9' {
			num := char - '0'
			out = out*10 + int(num)
		} else if char == '-' && i == 0 {
			sign = -1
		} else if char == '+' && i == 0 {
			sign = 1
		} else {
			return 0
		}
	}
	return out * sign
}

func main() {
	piscine.DescendComb()
	// a := SplitWhiteSpaces("Hello how are you?")
	// PrintWordsTables(a)

	// fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Wrong number of arguments")
// 		return
// 	}

// 	content, err := os.ReadFile(os.Args[1])
// 	if err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}

// 	fmt.Println(string(content))
// }
