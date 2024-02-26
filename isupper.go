package piscine

func IsUpper(s string) bool {
	flag := true

	for _, ch := range s {
		if ch >= 65 && ch <= 90 && flag {
			flag = true
		} else {
			flag = false
		}
	}
	if flag {
		return true
	} else {
		return false
	}
}
