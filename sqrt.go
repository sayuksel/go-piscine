package piscine

func Sqrt(nb int) int {
	sqrt := nb * nb
	if sqrt%2 == 0 {
		return sqrt
	} else {
		return 0
	}
}
