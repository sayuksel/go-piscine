package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	arr := []int{}
	remainder := 0

	for i := 0; i <= len(string(n+1)); i++ {
		remainder = n % 10
		n /= 10
		arr[i] = remainder

	}

	for i := 0; i <= len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = temp
		}
	}

	for i := 0; i <= len(arr); i++ {
		z01.PrintRune(rune(arr[i]))
	}
}
