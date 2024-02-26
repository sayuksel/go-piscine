package piscine

func IsLower(s string) bool {
	flag := true

	for _, ch := range s {
		if ch >= 97 && ch <= 122 && flag {
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
