package piscine

func SplitWhiteSpaces(s string) []string {
	arr := []rune(s)
	word := ""
	retarr := []string{}

	for i := 0; i < len(s); i++ {
		if arr[i] != ' ' && arr[i] != '\n' && arr[i] != '\t' {
			word += string(arr[i])
		} else if word != "" {
			retarr = append(retarr, word)
			word = ""
		}
	}
	retarr = append(retarr, word)
	return retarr
}
