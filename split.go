package piscine

import (
	"strings"
)

func splitString(s, sep string) []string {
	// Use strings.Split to split the string s by the separator sep
	return strings.Split(s, sep)
}

/*
func main() {
	// Test the function with a sample input string and separator
	s := "apple,banana,orange,grape"
	sep := ","
	result := splitString(s, sep)
	fmt.Println(result) // Output: [apple banana orange grape]
}
*/
