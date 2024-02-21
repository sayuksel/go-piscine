package piscine

func StrRev(s string) string {
	arr1 := []rune(s)
	arr2 := []rune{}
	for i := len(arr1) - 1; i >= 0; i-- {
		arr2 = append(arr2, arr1[i])
	}
	return string(arr2)
}
