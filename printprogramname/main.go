package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	input := os.Args

	line := []rune(input[0])

	for i := 0; i < len(line); i++ {
		if line[i] != '.' && line[i] != '/' {
			z01.PrintRune(line[i])
		}
	}
}
