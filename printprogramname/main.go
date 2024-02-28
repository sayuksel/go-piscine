package main

import (
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	
	fullName := args[0]
	trimmedName := fullName[2:]
	for _, v := range trimmedName {
		z01.PrintRune(v)
	}
	z01.PrintRune('\n')
}
