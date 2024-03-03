package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		filename := os.Args[1]
		file, error := os.Open(filename)

		if error != nil {
			fmt.Println(error)
		} else {
			data := make([]byte, 14)
			file.Read(data)
			fmt.Println(string(data))
			file.Close()
		}
	}
}
