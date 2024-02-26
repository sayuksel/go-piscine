package piscine

func ToLower(s string) string {
	arr := []rune(s)

	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			arr[i] += 32
		}
	}
	return string(arr)
}
