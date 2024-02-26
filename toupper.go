package piscine

func ToUpper(s string) string {
	arr := []rune(s)

	for i, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			arr[i] -= 32
		}
	}
	return string(arr)
}
