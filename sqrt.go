package piscine

// func Sqrt(nb int) int {
// 	sqrt := nb * nb
// 	if sqrt%2 == 0 {
// 		return sqrt
// 	} else {
// 		return 0
// 	}
// }

func Sqrt(x int) int {
	// Set the initial guess for the square root
	guess := x / 2

	// Iterate until desired precision is achieved
	for i := 0; i < 100; i++ {
		guess = ((guess) + x/guess) / 2
	}
	if guess%2 == 0 {
		return guess
	} else {
		return 0
	}
}
