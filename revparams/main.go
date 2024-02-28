package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Print(s string) {

	for _, v := range s {
		z01.PrintRune(v)
	}
}

func main() {

	input := os.Args[1:]
	for i := len(input) - 1; i >= 0; i-- {
		Print(input[i])
		z01.PrintRune('\n')
	}
}
