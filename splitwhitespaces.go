package piscine

func SplitWhiteSpaces(s string) []string {
	arr := []rune(s)
	retarr := []string{}
	lastr := 0
	count := 0

	for i, v := range arr {
		if v == ' ' || i == len(arr)-1 {
			retarr = append(retarr, string(arr[lastr:i]))
			lastr = i + 1
			count++
		}
	}
	return retarr
}
