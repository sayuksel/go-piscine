package piscine

func IsNumber(s rune) bool {
	if s >= '0' && s <= '9' {
		return true
	} else {
		return false
	}
}
