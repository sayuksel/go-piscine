// package piscine

// func IterativeFactorial(nb int) int {
// 	result := 1

// 	if nb <= 0 {
// 		return 0
// 	} else {
// 		for i := nb; i > 0; i-- {
// 			result *= i
// 		}
// 	}
// 	return result
// }

package piscine

func IterativeFactorial(nb int) int {
	result := 1

	if nb < 0 {
		return 0
	} else if nb == 0 {
		return 1
	} else {
		for i := nb; i > 0; i-- {
			result *= i
		}
	}
	return result
}
