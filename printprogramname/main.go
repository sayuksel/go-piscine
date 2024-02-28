package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	input := os.Args

	full := input[0]
	trimmed := full[2:]

	for _, v := range trimmed {
		z01.PrintRune(v)
	}
	// line := []rune(input[0])

	// for i := 0; i < len(line); i++ {
	// 	if line[i] != '.' && line[i] != '/' {
	// 		z01.PrintRune(line[i])
	// 	}
	// }
}
