package main

import "github.com/01-edu/z01"

func main() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				if i < j && j < k {
					if i != j && j != k {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(k)
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
				}
			}
		}
	}
}
