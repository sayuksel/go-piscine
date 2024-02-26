package piscine

func IsAlpha(s string) bool {
	for _, v := range s {
		if IsUpper(string(v)) || IsLower(string(v)) || IsNumber(v) {
			continue
		} else {
			return false
		}
	}
	return true
}
