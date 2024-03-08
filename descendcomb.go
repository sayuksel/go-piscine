package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for w := '9'; w >= '0'; w-- {
		for x := '9'; x >= '0'; x-- {
			for y := '9'; y >= '0'; y-- {
				for z := '9'; z >= '0'; z-- {
					z01.PrintRune(w)
					z01.PrintRune(x)
					z01.PrintRune(' ')
					z01.PrintRune(y)
					z01.PrintRune(z)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
