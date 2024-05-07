// https://stackoverflow.com/questions/7760545/escape-double-quotes-in-parameter
//'\!" #$%&'"'"'()*+,-./'
/*
__       _   _ _           _  _      _    _   __           _    __ __       _                                   __ $
\ \     | | ( | )        _| || |_   | |  (_) / /   ___    ( )  / / \ \   /\| |/\     _                         / / $
 \ \    | |  V V        |_  __  _| / __)    / /   ( _ )   |/  | |   | |  \ ` ' /   _| |_       ______         / /  $
  \ \   | |              _| || |_  \__ \   / /    / _ \/\     | |   | | |_     _| |_   _|     |______|       / /   $
   \ \  |_|             |_  __  _| (   /  / / _  | (_>  <     | |   | |  / , . \    |_|    _            _   / /    $
    \_\ (_)               |_||_|    |_|  /_/ (_)  \___/\/     | |   | |  \/|_|\/          ( )          (_) /_/     $
                                                               \_\ /_/                    |/                       $
                                                                                                                   $
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Wrong number of arguments. os.Args length: ", len(os.Args))
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
		case '!':
			if i != len(arr)-1 {
				if arr[i+1] == '"' {
					line1 += " _  "
					line2 += "| | "
					line3 += "| | "
					line4 += "| | "
					line5 += "|_| "
					line6 += "(_) "
					line7 += "    "
					line8 += "    "

					line1 += " _ _ "
					line2 += "( | )"
					line3 += " V V "
					line4 += "     "
					line5 += "     "
					line6 += "     "
					line7 += "     "
					line8 += "     "

					i++
				} else {
					line1 += " _  "
					line2 += "| | "
					line3 += "| | "
					line4 += "| | "
					line5 += "|_| "
					line6 += "(_) "
					line7 += "    "
					line8 += "    "
				}
			} else {
				line1 += " _  "
				line2 += "| | "
				line3 += "| | "
				line4 += "| | "
				line5 += "|_| "
				line6 += "(_) "
				line7 += "    "
				line8 += "    "
			}

		case '#':
			line1 += "   _  _    "
			line2 += " _| || |_  "
			line3 += "|_  __  _| "
			line4 += " _| || |_  "
			line5 += "|_  __  _| "
			line6 += "  |_||_|   "
			line7 += "           "
			line8 += "           "

		case '$':
			line1 += "  _   "
			line2 += " | |  "
			line3 += "/ __) "
			line4 += "\\__ \\ "
			line5 += "(   / "
			line6 += " |_|  "
			line7 += "      "
			line8 += "      "

		case '%':
			line1 += " _   __ "
			line2 += "(_) / / "
			line3 += "   / /  "
			line4 += "  / /   "
			line5 += " / / _  "
			line6 += "/_/ (_) "
			line7 += "        "
			line8 += "        "

		case '&':
			line1 += "         "
			line2 += "  ___    "
			line3 += " ( _ )   "
			line4 += " / _ \\/\\ "
			line5 += "| (_>  < "
			line6 += " \\___/\\/ "
			line7 += "         "
			line8 += "         "

		case '\'':
			line1 += " _  "
			line2 += "( ) "
			line3 += "|/  "
			line4 += "    "
			line5 += "    "
			line6 += "    "
			line7 += "    "
			line8 += "    "

		case '(':
			line1 += "  __ "
			line2 += " / / "
			line3 += "| |  "
			line4 += "| |  "
			line5 += "| |  "
			line6 += "| |  "
			line7 += " \\_\\ "
			line8 += "     "

		case ')':
			line1 += "__   "
			line2 += "\\ \\  "
			line3 += " | | "
			line4 += " | | "
			line5 += " | | "
			line6 += " | | "
			line7 += "/_/  "
			line8 += "     "

		case '*':
			line1 += "    _     "
			line2 += " /\\| |/\\  "
			line3 += " \\ ` ' /  "
			line4 += "|_     _| "
			line5 += " / , . \\  "
			line6 += " \\/|_|\\/  "
			line7 += "          "
			line8 += "          "

		case '+':
			line1 += "        "
			line2 += "   _    "
			line3 += " _| |_  "
			line4 += "|_   _| "
			line5 += "  |_|   "
			line6 += "        "
			line7 += "        "
			line8 += "        "

		case ',':
			line1 += "    "
			line2 += "    "
			line3 += "    "
			line4 += "    "
			line5 += " _  "
			line6 += "( ) "
			line7 += "|/  "
			line8 += "    "

		case '-':
			line1 += "         "
			line2 += "         "
			line3 += " ______  "
			line4 += "|______| "
			line5 += "         "
			line6 += "         "
			line7 += "         "
			line8 += "         "

		case '.':
			line1 += "    "
			line2 += "    "
			line3 += "    "
			line4 += "    "
			line5 += " _  "
			line6 += "(_) "
			line7 += "    "
			line8 += "    "

		case '/':
			line1 += "     __ "
			line2 += "    / / "
			line3 += "   / /  "
			line4 += "  / /   "
			line5 += " / /    "
			line6 += "/_/     "
			line7 += "        "
			line8 += "        "

		case '0':
			line1 += "  ___   "
			line2 += " / _ \\  "
			line3 += "| | | | "
			line4 += "| |_| | "
			line5 += " \\___/  "
			line6 += "        "
			line7 += "        "
			line8 += "        "

		case '1':
			line1 += "    "
			line2 += " _  "
			line3 += "/ | "
			line4 += "| | "
			line5 += "| | "
			line6 += "|_| "
			line7 += "    "
			line8 += "    "

		case '2':
			line1 += "        "
			line2 += " ____   "
			line3 += "|___ \\  "
			line4 += "  __) | "
			line5 += " / __/  "
			line6 += "|_____| "
			line7 += "        "
			line8 += "        "

		case '3':
			line1 += "        "
			line2 += " _____  "
			line3 += "|___ /  "
			line4 += "  |_ \\  "
			line5 += " ___) | "
			line6 += "|____/  "
			line7 += "        "
			line8 += "        "

		case '4':
			line1 += "         "
			line2 += " _  _    "
			line3 += "| || |   "
			line4 += "| || |_  "
			line5 += "|__   _| "
			line6 += "   |_|   "
			line7 += "         "
			line8 += "         "
		case '5':
			line1 += "        "
			line2 += " ____   "
			line3 += "| ___|  "
			line4 += "|___ \\  "
			line5 += "  __) | "
			line6 += "|____/  "
			line7 += "        "
			line8 += "        "
		case '6':
			line1 += "        "
			line2 += "  __    "
			line3 += " / /    "
			line4 += "| '_ \\  "
			line5 += "| (_) | "
			line6 += " \\___/  "
			line7 += "        "
			line8 += "        "
		case '7':
			line1 += "        "
			line2 += " _____  "
			line3 += "|___  | "
			line4 += "   / /  "
			line5 += "  / /   "
			line6 += " /_/    "
			line7 += "        "
			line8 += "        "
		case '8':
			line1 += "        "
			line2 += "  ___   "
			line3 += " ( _ )  "
			line4 += " / _ \\  "
			line5 += "| (_) | "
			line6 += " \\___/  "
			line7 += "        "
			line8 += "        "
		case '9':
			line1 += "        "
			line2 += "  ___   "
			line3 += " / _ \\  "
			line4 += "| (_) | "
			line5 += " \\__, | "
			line6 += "   / /  "
			line7 += "  /_/   "
			line8 += "        "
		case ':':
			line1 += "    "
			line2 += " _  "
			line3 += "(_) "
			line4 += "    "
			line5 += " _  "
			line6 += "(_) "
			line7 += "    "
			line8 += "    "
		case ';':
			line1 += "    "
			line2 += " _  "
			line3 += "(_) "
			line4 += "    "
			line5 += " _  "
			line6 += "( ) "
			line7 += "|/  "
			line8 += "    "
		case '<':
			line1 += "   __ "
			line2 += "  / / "
			line3 += " / /  "
			line4 += "< <   "
			line5 += " \\ \\  "
			line6 += "  \\_\\ "
			line7 += "      "
			line8 += "      "
		case '=':
			line1 += "         "
			line2 += " ______  "
			line3 += "|______| "
			line4 += " ______  "
			line5 += "|______| "
			line6 += "         "
			line7 += "         "
			line8 += "         "
		case '>':
			line1 += "__    "
			line2 += "\\ \\   "
			line3 += " \\ \\  "
			line4 += "  > > "
			line5 += " / /  "
			line6 += "/_/   "
			line7 += "      "
			line8 += "      "
		case '?':
			line1 += " ___   "
			line2 += "|__ \\  "
			line3 += "   ) | "
			line4 += "  / /  "
			line5 += " |_|   "
			line6 += " (_)   "
			line7 += "       "
			line8 += "       "
		case '@':
			line1 += "   ____   "
			line2 += "  / __ \\  "
			line3 += " / / _` | "
			line4 += "| | (_| | "
			line5 += " \\ \\__,_| "
			line6 += "  \\____/  "
			line7 += "          "
			line8 += "          "
		case '[':
			line1 += " ___  "
			line2 += "|  _| "
			line3 += "| |   "
			line4 += "| |   "
			line5 += "| |   "
			line6 += "| |_  "
			line7 += "|___| "
			line8 += "      "
		case ']':
			line1 += " ___  "
			line2 += "|_  | "
			line3 += "  | | "
			line4 += "  | | "
			line5 += "  | | "
			line6 += " _| | "
			line7 += "|___| "
			line8 += "      "
		case '^':
			line1 += " /\\  "
			line2 += "|/\\| "
			line3 += "     "
			line4 += "     "
			line5 += "     "
			line6 += "     "
			line7 += "     "
			line8 += "     "
		case '_':
			line1 += "         "
			line2 += "         "
			line3 += "         "
			line4 += "         "
			line5 += "         "
			line6 += "         "
			line7 += " ______  "
			line8 += "|______| "
		case '`':
			line1 += " _  "
			line2 += "( ) "
			line3 += " \\| "
			line4 += "    "
			line5 += "    "
			line6 += "    "
			line7 += "    "
			line8 += "    "
		case '{':
			line1 += "   __ "
			line2 += "  / / "
			line3 += " | |  "
			line4 += "/ /   "
			line5 += "\\ \\   "
			line6 += " | |  "
			line7 += "  \\_\\ "
			line8 += "      "
		case '|':
			line1 += " _  "
			line2 += "| | "
			line3 += "| | "
			line4 += "| | "
			line5 += "| | "
			line6 += "| | "
			line7 += "| | "
			line8 += "|_| "
		case '}':
			line1 += "__    "
			line2 += "\\ \\   "
			line3 += " | |  "
			line4 += "  \\ \\ "
			line5 += "  / / "
			line6 += " | |  "
			line7 += "/_/   "
			line8 += "      "
		case '~':
			line1 += " /\\/| "
			line2 += "|/\\/  "
			line3 += "      "
			line4 += "      "
			line5 += "      "
			line6 += "      "
			line7 += "      "
			line8 += "      "
		case 'a':
			line1 += "        "
			line2 += "        "
			line3 += "  __ _  "
			line4 += " / _` | "
			line5 += "| (_| | "
			line6 += " \\__,_| "
			line7 += "        "
			line8 += "        "
		case 'b':
			line1 += " _      "
			line2 += "| |     "
			line3 += "| |__   "
			line4 += "| '_ \\  "
			line5 += "| |_) | "
			line6 += "|_.__/  "
			line7 += "        "
			line8 += "        "
		case 'c':
			line1 += "       "
			line2 += "       "
			line3 += "  ___  "
			line4 += " / __| "
			line5 += "| (__  "
			line6 += " \\___| "
			line7 += "       "
			line8 += "       "
		case 'd':
			line1 += "     _  "
			line2 += "    | | "
			line3 += "  __| | "
			line4 += " / _` | "
			line5 += "| (_| | "
			line6 += " \\__,_| "
			line7 += "        "
			line8 += "        "
		case 'e':
			line1 += "       "
			line2 += "       "
			line3 += "  ___  "
			line4 += " / _ \\ "
			line5 += "|  __/ "
			line6 += " \\___| "
			line7 += "       "
			line8 += "       "
		case 'f':
			line1 += "  __  "
			line2 += " / _| "
			line3 += "| |_  "
			line4 += "|  _| "
			line5 += "| |   "
			line6 += "|_|   "
			line7 += "      "
			line8 += "      "
		case 'g':
			line1 += "        "
			line2 += "        "
			line3 += "  __ _  "
			line4 += " / _` | "
			line5 += "| (_| | "
			line6 += " \\__, | "
			line7 += "  __/ | "
			line8 += " |___/  "
		case 'h':
			line1 += " _      "
			line2 += "| |     "
			line3 += "| |__   "
			line4 += "|  _ \\  "
			line5 += "| | | | "
			line6 += "|_| |_| "
			line7 += "        "
			line8 += "        "
		case 'i':
			line1 += " _  "
			line2 += "(_) "
			line3 += " _  "
			line4 += "| | "
			line5 += "| | "
			line6 += "|_| "
			line7 += "    "
			line8 += "    "
		case 'j':
			line1 += "   _  "
			line2 += "  (_) "
			line3 += "   _  "
			line4 += "  | | "
			line5 += "  | | "
			line6 += "  | | "
			line7 += " _/ | "
			line8 += "|__/  "
		case 'k':
			line1 += "       "
			line2 += " _     "
			line3 += "| | _  "
			line4 += "| |/ / "
			line5 += "|   <  "
			line6 += "|_|\\_\\ "
			line7 += "       "
			line8 += "       "
		case 'l':
			line1 += " _  "
			line2 += "| | "
			line3 += "| | "
			line4 += "| | "
			line5 += "| | "
			line6 += "|_| "
			line7 += "    "
			line8 += "    "

		case 'm':
			line1 += "            "
			line2 += "            "
			line3 += " _ __ ___   "
			line4 += "| '_ ` _ \\  "
			line5 += "| | | | | | "
			line6 += "|_| |_| |_| "
			line7 += "            "
			line8 += "            "

		case 'n':
			line1 += "        "
			line2 += "        "
			line3 += " _ __   "
			line4 += "| '_ \\  "
			line5 += "| | | | "
			line6 += "|_| |_| "
			line7 += "        "
			line8 += "        "

		case 'o':
			line1 += "        "
			line2 += "        "
			line3 += "  ___   "
			line4 += " / _ \\  "
			line5 += "| (_) | "
			line6 += " \\___/  "
			line7 += "        "
			line8 += "        "

		case 'p':
			line1 += "        "
			line2 += "        "
			line3 += " _ __   "
			line4 += "| '_ \\  "
			line5 += "| |_) | "
			line6 += "| .__/  "
			line7 += "| |     "
			line8 += "|_|     "

		case 'q':
			line1 += "        "
			line2 += "        "
			line3 += "  __ _  "
			line4 += " / _` | "
			line5 += "| (_| | "
			line6 += " \\__, | "
			line7 += "    | | "
			line8 += "    |_| "
		case 'r':
			line1 += "       "
			line2 += "       "
			line3 += " _ __  "
			line4 += "| '__| "
			line5 += "| |    "
			line6 += "|_|    "
			line7 += "       "
			line8 += "       "
		case 's':
			line1 += "      "
			line2 += "      "
			line3 += " ___  "
			line4 += "/ __| "
			line5 += "\\__ \\ "
			line6 += "|___/ "
			line7 += "      "
			line8 += "      "
		case 't':
			line1 += " _    "
			line2 += "| |   "
			line3 += "| |_  "
			line4 += "| __| "
			line5 += "\\ |_  "
			line6 += " \\__| "
			line7 += "      "
			line8 += "      "
		case 'u':
			line1 += "        "
			line2 += "        "
			line3 += " _   _  "
			line4 += "| | | | "
			line5 += "| |_| | "
			line6 += " \\__,_| "
			line7 += "        "
			line8 += "        "
		case 'v':
			line1 += "        "
			line2 += "        "
			line3 += "__   __ "
			line4 += "\\ \\ / / "
			line5 += " \\ V /  "
			line6 += "  \\_/   "
			line7 += "        "
			line8 += "        "
		case 'w':
			line1 += "           "
			line2 += "           "
			line3 += "__      __ "
			line4 += "\\ \\ /\\ / / "
			line5 += " \\ V  V /  "
			line6 += "  \\_/\\_/   "
			line7 += "           "
			line8 += "           "
		case 'x':
			line1 += "       "
			line2 += "       "
			line3 += "__  __ "
			line4 += "\\ \\/ / "
			line5 += " >  <  "
			line6 += "/_/\\_\\ "
			line7 += "       "
			line8 += "       "
		case 'y':
			line1 += "        "
			line2 += "        "
			line3 += " _   _  "
			line4 += "| | | | "
			line5 += "| |_| | "
			line6 += " \\__, | "
			line7 += " __/ /  "
			line8 += "|___/   "
		case 'z':
			line1 += "      "
			line2 += "      "
			line3 += " ____ "
			line4 += "|_  / "
			line5 += " / /  "
			line6 += "/___| "
			line7 += "      "
			line8 += "      "

		case 'A':
			line1 += "           "
			line2 += "    /\\     "
			line3 += "   /  \\    "
			line4 += "  / /\\ \\   "
			line5 += " / ____ \\  "
			line6 += "/_/    \\_\\ "
			line7 += "           "
			line8 += "           "
		case 'B':
			line1 += " ____   "
			line2 += "|  _ \\  "
			line3 += "| |_) | "
			line4 += "|  _ <  "
			line5 += "| |_) | "
			line6 += "|____/  "
			line7 += "        "
			line8 += "        "
		case 'C':
			line1 += "  _____  "
			line2 += " / ____| "
			line3 += "| |      "
			line4 += "| |      "
			line5 += "| |____  "
			line6 += " \\_____| "
			line7 += "         "
			line8 += "         "
		case 'D':
			line1 += " _____   "
			line2 += "|  __ \\  "
			line3 += "| |  | | "
			line4 += "| |  | | "
			line5 += "| |__| | "
			line6 += "|_____/  "
			line7 += "         "
			line8 += "         "
		case 'E':
			line1 += " ______  "
			line2 += "|  ____| "
			line3 += "| |__    "
			line4 += "|  __|   "
			line5 += "| |____  "
			line6 += "|______| "
			line7 += "         "
			line8 += "         "
		case 'F':
			line1 += " ______  "
			line2 += "|  ____| "
			line3 += "| |__    "
			line4 += "|  __|   "
			line5 += "| |      "
			line6 += "|_|      "
			line7 += "         "
			line8 += "         "
		case 'G':
			line1 += "  _____  "
			line2 += " / ____| "
			line3 += "| |  __  "
			line4 += "| | |_ | "
			line5 += "| |__| | "
			line6 += " \\_____| "
			line7 += "         "
			line8 += "         "
		case 'H':
			line1 += " _    _  "
			line2 += "| |  | | "
			line3 += "| |__| | "
			line4 += "|  __  | "
			line5 += "| |  | | "
			line6 += "|_|  |_| "
			line7 += "         "
			line8 += "         "
		case 'I':
			line1 += " _____  "
			line2 += "|_   _| "
			line3 += "  | |   "
			line4 += "  | |   "
			line5 += " _| |_  "
			line6 += "|_____| "
			line7 += "        "
			line8 += "        "
		case 'J':
			line1 += "      _  "
			line2 += "     | | "
			line3 += "     | | "
			line4 += " _   | | "
			line5 += "| |__| | "
			line6 += " \\____/  "
			line7 += "         "
			line8 += "         "
		case 'K':
			line1 += " _  __ "
			line2 += "| |/ / "
			line3 += "| ' /  "
			line4 += "|  <   "
			line5 += "| . \\  "
			line6 += "|_|\\_\\ "
			line7 += "       "
			line8 += "       "
		case 'L':
			line1 += " _       "
			line2 += "| |      "
			line3 += "| |      "
			line4 += "| |      "
			line5 += "| |____  "
			line6 += "|______| "
			line7 += "         "
			line8 += "         "
		case 'M':
			line1 += " __  __  "
			line2 += "|  \\/  | "
			line3 += "| \\  / | "
			line4 += "| |\\/| | "
			line5 += "| |  | | "
			line6 += "|_|  |_| "
			line7 += "         "
			line8 += "         "
		case 'N':
			line1 += " _   _  "
			line2 += "| \\ | | "
			line3 += "|  \\| | "
			line4 += "| . ` | "
			line5 += "| |\\  | "
			line6 += "|_| \\_| "
			line7 += "        "
			line8 += "        "
		case 'O':
			line1 += "  ____   "
			line2 += " / __ \\  "
			line3 += "| |  | | "
			line4 += "| |  | | "
			line5 += "| |__| | "
			line6 += " \\____/  "
			line7 += "         "
			line8 += "         "
		case 'P':
			line1 += " _____   "
			line2 += "|  __ \\  "
			line3 += "| |__) | "
			line4 += "|  ___/  "
			line5 += "| |      "
			line6 += "|_|      "
			line7 += "         "
			line8 += "         "
		case 'Q':
			line1 += "  ____   "
			line2 += " / __ \\  "
			line3 += "| |  | | "
			line4 += "| |  | | "
			line5 += "| |__| | "
			line6 += " \\___\\_\\ "
			line7 += "         "
			line8 += "         "
		case 'R':
			line1 += " _____   "
			line2 += "|  __ \\  "
			line3 += "| |__) | "
			line4 += "|  _  /  "
			line5 += "| | \\ \\  "
			line6 += "|_|  \\_\\ "
			line7 += "         "
			line8 += "         "
		case 'S':
			line1 += "  _____  "
			line2 += " / ____| "
			line3 += "| (___   "
			line4 += " \\___ \\  "
			line5 += " ____) | "
			line6 += "|_____/  "
			line7 += "         "
			line8 += "         "
		case 'T':
			line1 += " _______  "
			line2 += "|__   __| "
			line3 += "   | |    "
			line4 += "   | |    "
			line5 += "   | |    "
			line6 += "   |_|    "
			line7 += "          "
			line8 += "          "
		case 'U':
			line1 += " _    _  "
			line2 += "| |  | | "
			line3 += "| |  | | "
			line4 += "| |  | | "
			line5 += "| |__| | "
			line6 += " \\____/  "
			line7 += "         "
			line8 += "         "
		case 'V':
			line1 += "__      __ "
			line2 += "\\ \\    / / "
			line3 += " \\ \\  / /  "
			line4 += "  \\ \\/ /   "
			line5 += "   \\  /    "
			line6 += "    \\/     "
			line7 += "           "
			line8 += "           "
		case 'W':
			line1 += "__          __ "
			line2 += "\\ \\        / / "
			line3 += " \\ \\  /\\  / /  "
			line4 += "  \\ \\/  \\/ /   "
			line5 += "   \\  /\\  /    "
			line6 += "    \\/  \\/     "
			line7 += "               "
			line8 += "               "
		case 'X':
			line1 += "__   __ "
			line2 += "\\ \\ / / "
			line3 += " \\ V /  "
			line4 += "  > <   "
			line5 += " / . \\  "
			line6 += "/_/ \\_\\ "
			line7 += "        "
			line8 += "        "
		case 'Y':
			line1 += "__     __ "
			line2 += "\\ \\   / / "
			line3 += " \\ \\_/ /  "
			line4 += "  \\   /   "
			line5 += "   | |    "
			line6 += "   |_|    "
			line7 += "          "
			line8 += "          "
		case 'Z':
			line1 += " ______ "
			line2 += "|___  / "
			line3 += "   / /  "
			line4 += "  / /   "
			line5 += " / /__  "
			line6 += "/_____| "
			line7 += "        "
			line8 += "        "

		case '"':
			line1 += " _ _  "
			line2 += "( | ) "
			line3 += " V V  "
			line4 += "      "
			line5 += "      "
			line6 += "      "
			line7 += "      "
			line8 += "      "

		case '\\':
			if i != len(arr)-1 {
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
				} else if arr[i+1] == '"' {
					line1 += " _ _ "
					line2 += "( | )"
					line3 += " V V "
					line4 += "     "
					line5 += "     "
					line6 += "     "
					line7 += "     "
					line8 += "     "
				} else {
					line1 += " __      "
					line2 += " \\ \\     "
					line3 += "  \\ \\    "
					line4 += "   \\ \\   "
					line5 += "    \\ \\  "
					line6 += "     \\_\\ "
					line7 += "         "
					line8 += "         "
				}
			} else {
				line1 += " __      "
				line2 += " \\ \\     "
				line3 += "  \\ \\    "
				line4 += "   \\ \\   "
				line5 += "    \\ \\  "
				line6 += "     \\_\\ "
				line7 += "         "
				line8 += "         "
			}

		case ' ':
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
	if line1 != "" {
		fmt.Println(line1)
		fmt.Println(line2)
		fmt.Println(line3)
		fmt.Println(line4)
		fmt.Println(line5)
		fmt.Println(line6)
		fmt.Println(line7)
		fmt.Println(line8)
	}
}
