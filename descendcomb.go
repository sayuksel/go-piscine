// package piscine

// import "github.com/01-edu/z01"

// func DescendComb() {
// 	for w := '9'; w >= '0'; w-- {
// 		for x := '9'; x >= '0'; x-- {
// 			for y := '9'; y >= '0'; y-- {
// 				for z := '9'; z >= '0'; z-- {
// 					if w != y && x != z {
// 						z01.PrintRune(w)
// 						z01.PrintRune(x)
// 						z01.PrintRune(' ')
// 						z01.PrintRune(y)
// 						z01.PrintRune(z)
// 						z01.PrintRune(',')
// 						z01.PrintRune(' ')
// 					}
// 				}
// 			}
// 		}
// 	}
// }

package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 99; i >= 1; i-- {
		for j := i - 1; j >= 0; j-- {
			if i != j {
				z01.PrintRune(rune(i/10 + 48))
				z01.PrintRune(rune(i%10 + 48))
				z01.PrintRune(' ')
				z01.PrintRune(rune(j/10 + 48))
				z01.PrintRune(rune(j%10 + 48))
				if i != 1 {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
