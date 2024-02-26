package piscine

func IsPrintable(s string) bool {
	for _, ch := range s {
		if ch < 32 || ch >= 127 {
			return false
		}
	}
	return true
}
