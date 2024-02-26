package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)

	// 	if n > len(arr)-1 || n < 0 {
	// 		return 0
	// 	} else {
	// 		return arr[n+1]
	// 	}
	// }

	// for i:= 0; i < len(arr) -1; i++ {
	// 	if arr[i] == rune(n) {
	// 		return arr[i]
	// 	}
	// }
	// return 0

	// for index, ch := range arr {
	// 	if index -1  == n {
	// 		return ch
	// 	}
	// }
	// return 0
	if n <= 0 || len(arr) < n {
		return 0
	}

	return arr[n-1]
}
