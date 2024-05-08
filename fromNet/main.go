/*package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		return
	}

	// opening the file in read-only mode. The file must exist (in the current working directory)
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	asciiChars := make(map[int][]string)

	dec := 31

	for _, line := range lines {
		if line == "" {
			dec++
		} else {
			asciiChars[dec] = append(asciiChars[dec], line)
		}
	}

	// string that user inputs
	userInput := os.Args[1]
	Newline(userInput, asciiChars)
}

// Newline function prints string horizontally and with new line if user specifies
func Newline(n string, asciiChars map[int][]string) {
	var currentLine string
	var precedingChar rune

	for _, char := range n {
		if char == '\n' && (precedingChar == '\n' || precedingChar == 0) {
			for k := 0; k < 8; k++ {
				fmt.Println() // print 8 new lines
			}
			currentLine = "" // reset line buffer after printing new lines
		} else if char == '\n' {
			fmt.Println(currentLine) // print the accumulated line
			currentLine = ""          // reset line buffer after printing
		} else {
			currentLine += asciiLine(char, asciiChars) // accumulate the ASCII art line
		}
		precedingChar = char
	}

	// Print the accumulated line if there are remaining characters
	if currentLine != "" {
		fmt.Println(currentLine)
	}
}

// asciiLine function returns the ASCII art lines for a given character
func asciiLine(char rune, asciiChars map[int][]string) string {
	asciiArt := asciiChars[int(char)]
	var lines string
	for _, line := range asciiArt {
		lines += line // accumulate ASCII art lines
	}
	return lines
}*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments!")
		return
	}

	// to check values of each argument in the passed arguments
	// for i := 0; i < len(os.Args); i++ {
	// 	fmt.Println("index", i, " : ", os.Args[i])
	// }

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
	userInput := os.Args[1]
	Newline(userInput, asciiChrs)
}

// Newline function prints string horizontally and with new line if user specifies
func Newline(n string, y map[int][]string) {
	replaceNewline := strings.ReplaceAll(n, "\\n", "\n")
	wordsSlice := strings.Split(replaceNewline, "\n")
	counter := 0
	for _, word := range wordsSlice {
		flag := false
		for j := 0; j < len(y[32]); j++ {
			for _, letter := range word {
					fmt.Print(y[int(letter)][j])
			}
			
if len(word) != 0{
	fmt.Println("")
} else if !flag && counter > 0{
	fmt.Println("")
	flag = true
}
						
					
		
			
			
		}
		counter++
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