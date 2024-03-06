package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fileName := args[1]

	if len(args) > 2 {
		fmt.Println("too many arguments")
	}

	if len(args) == 1 {
		fmt.Println("File name missing")
	}

	file, error := os.Open(fileName)

	if error != nil {
		fmt.Println(error)
	} else {
		data := make([]byte, 14)
		file.Read(data)
		fmt.Println(string(data))
		file.Close()
	}
}
