package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	a := ' '
	setPoint(points)

	for i := 0; i < 14; i++ {
		switch i {
		case 0:
			a = 'x'
		case 1:
			a = ' '
		case 2:
			a = '='
		case 3:
			a = ' '
		case 4:
			a = 52
		case 5:
			a = 50
		case 6:
			a = ','
		case 7:
			a = ' '
		case 8:
			a = 'y'
		case 9:
			a = ' '
		case 10:
			a = '='
		case 11:
			a = ' '
		case 12:
			a = 50
		case 13:
			a = 49

		}
		z01.PrintRune(a)
	}
	z01.PrintRune('\n')
}
