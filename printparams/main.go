package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	input := os.Args
	for _, x := range input {
		for _, y := range x {
			z01.PrintRune(y)
		}
		z01.PrintRune('\n')
	}
}
