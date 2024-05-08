package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2{
		fmt.Println("Wrong number of arguments!")
		return
	}
	for i :=0; i< len(os.Args); i++ {
		fmt.Println("index" , i, " : ", os.Args[i])
	}
	

	// opening the file in read-only mode. The file must exist (in the current working directory)
	file, _ := os.Open("standard.txt")

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	scanned := bufio.NewScanner(file)

	scanned.Split(bufio.ScanLines)

	var lines []string

	for scanned.Scan() {
		lines = append(lines, scanned.Text())
		
	}

	file.Close()

	asciiChrs := make(map[int][]string)

	dec := 31

	for _, line := range lines {
		if line == "" {
			dec++
		} else {
			asciiChrs[dec] = append(asciiChrs[dec], line)
		}
	}
	// string that user inputs
	if (len(os.Args[1]) > 1){
		userInput := os.Args[1]
		Newline(userInput, asciiChrs)
	} else {
		fmt.Println("Wrong Input!")
	}
	
	
}

// Newline function prints string horizontally and with new line if user specifies
func Newline(n string, y map[int][]string) {
	replaceNewline := strings.ReplaceAll(n, "\\n", "\n")
	wordsSlice := strings.Split(replaceNewline, "\n")
	for _, word := range wordsSlice {
		for j := 0; j < len(y[32]); j++ {
			for _, letter := range word {
				
				if y[int(letter)][j] == string('\n'){
					fmt.Println()
					break
			} else {
				fmt.Print(y[int(letter)][j])
			}
			}
			fmt.Println()
		}
	}
}

/*func Newline(n string, y map[int][]string) {
	replaceNewline := strings.ReplaceAll(n, "\\n", "\n")
	wordsSlice := strings.Split(replaceNewline, "\n")
	if wordsSlice[0] == "" {
		// If there are no characters before \n, print only one new line
		fmt.Println()
	} else {
		// If there are characters before \n, print eight new lines
		for i := 0; i < 8; i++ {
			fmt.Println()
		}
	}
	for _, word := range wordsSlice {
		for j := 0; j < len(y[32]); j++ {
			for _, letter := range word {
				fmt.Print(y[int(letter)][j])
			}
			fmt.Println()
		}
	}
}*/

