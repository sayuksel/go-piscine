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
			a = '4'
		case 5:
			a = '2'
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
			a = '2'
		case 13:
			a = '1'

		}
		z01.PrintRune(a)
	}
}
