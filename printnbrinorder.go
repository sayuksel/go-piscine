package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {

	if n == 0 {
		z01.PrintRune('0')
	}

	count := 0
	temp := n
	for temp > 0 {
		temp /= 10
		count++
	}

	// // Count the number of digits

	arr := make([]int, count)
	for i := count - 1; i >= 0; i-- {
		arr[i] = n % 10
		n /= 10
	}

	// // Extract the digits into an array

	for i := 0; i < count-1; i++ {
		for j := 0; j < count-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	// // Sort the digits in ascending order using bubble sort
	// for i := 0; i < count-1; i++ {
	// 	for j := 0; j < count-i-1; j++ {
	// 		if digits[j] > digits[j+1] {
	// 			digits[j], digits[j+1] = digits[j+1], digits[j]
	// 		}
	// 	}
	// }

	for _, ch := range arr {
		z01.PrintRune(rune(ch) + 48)
	}
}

// // Print the sorted digits
// for _, digit := range digits {
// 	z01.PrintRune(rune(digit)+'0')
// }
