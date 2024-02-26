package piscine

func IsUpper(s string) bool {
	flag := false
	for _, ch := range s {
		flag = false
		if ch >= 'A' && ch <= 'Z' {
			flag = true
		}
	}
	if flag {
		return true
	} else {
		return false
	}
}
