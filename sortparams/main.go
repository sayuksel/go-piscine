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

	for i := 0; i <= len(input)-1; i++ {
		for j := 0; j <= len(input)-i-1; j++ {
			if input[i] > input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
	}
	for _, v := range input {
		Print(v)
		z01.PrintRune('\n')
	}
}
