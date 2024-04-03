package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments")
		return
	}

	line1 := ""
	line2 := ""
	line3 := ""
	line4 := ""
	line5 := ""
	line6 := ""
	line7 := ""
	line8 := ""

	arr := []rune(os.Args[1])
	for i := 0; i < len(arr); i++ {
		switch arr[i] {

		case 'e':
			line1 += "       "
			line2 += "  ___  "
			line3 += " / _ \\ "
			line4 += "|  __/ "
			line5 += " \\___| "
			line6 += "       "
			line7 += "       "
			line8 += "       "

		case 'h':
			line1 += " _      "
			line2 += "| |     "
			line3 += "| |__   "
			line4 += "|  _ \\  "
			line5 += "| | | | "
			line6 += "|_| |_| "
			line7 += "        "
			line8 += "        "

		case 'l':
			line1 += " _  "
			line2 += "| | "
			line3 += "| | "
			line4 += "| | "
			line5 += "| | "
			line6 += "|_| "
			line7 += "    "
			line8 += "    "

		case 'o':
			line1 += "        "
			line2 += "  ___   "
			line3 += " / _ \\  "
			line4 += "| (_) | "
			line5 += " \\___/  "
			line6 += "        "
			line7 += "        "
			line8 += "        "

		case '\\':
			if arr[i] != arr[len(arr)-1] {
				if arr[i+1] == 'n' {
					if line1 != "" {
						fmt.Println(line1)
						fmt.Println(line2)
						fmt.Println(line3)
						fmt.Println(line4)
						fmt.Println(line5)
						fmt.Println(line6)
						fmt.Println(line7)
						fmt.Println(line8)

						line1 = ""
						line2 = ""
						line3 = ""
						line4 = ""
						line5 = ""
						line6 = ""
						line7 = ""
						line8 = ""
						i++
					} else {
						fmt.Println()
						i++
					}

				}

			} else {

			}

		default:
			line1 += "     "
			line2 += "     "
			line3 += "     "
			line4 += "     "
			line5 += "     "
			line6 += "     "
			line7 += "     "
			line8 += "     "

		}
	}
	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line3)
	fmt.Println(line4)
	fmt.Println(line5)
	fmt.Println(line6)
	fmt.Println(line7)
	fmt.Println(line8)
}
