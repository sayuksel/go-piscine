package piscine

func AlphaCount(s string) int {
	alpha := 0

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			alpha++
		} else if ch >= 'A' && ch <= 'Z' {
			alpha++
		}
	}
	return alpha
}
