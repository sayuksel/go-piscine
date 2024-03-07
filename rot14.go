package piscine

func Rot14(s string) string {
	arr := []rune(s)
	newarr := ""
	// lower := false
	// upper := false

	for _, v := range arr {
		if v >= 'A' && v <= 'Z' {
			if v >= 'M' {
				newarr += string(v - 12)
			} else {
				newarr += string(v + 14)
			}
		} else if v >= 'a' && v <= 'z' {
			if v >= 'm' {
				newarr += string(v - 12)
			} else {
				newarr += string(v + 14)
			}
		} else {
			newarr += string(v)
		}
	}
	return newarr
}
