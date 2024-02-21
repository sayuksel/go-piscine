package piscine

/*func BasicAtoi(s string) {
	arr_s := []rune(s)
	arr_int := []int{}

	for i := 0; i < len(arr_s)-1; i++ {
		arr_int = append(arr_int, int(arr_s[i]))
	}
	int_s := arr_int
	return int_s
}
*/

func BasicAtoi(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		result = (result * 10) + int(s[i]-48)
	}
	return result
}
