package piscine

func AlphaCount(s string) int {
	alpha := 0

	for _, ch := range s {
		if ch >= 97 && ch <= 122 {
			alpha++
		} else if ch >= 65 && ch <= 90 {
			alpha++
		}
	}
	return alpha
}
