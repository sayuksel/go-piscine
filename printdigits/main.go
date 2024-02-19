package main

import "github.com/01-edu/z01"

func main() {
	for i := '0'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')

	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)

}
