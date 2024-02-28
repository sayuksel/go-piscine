package piscine

func Capitalize(s string) string {
	IsFirst := true
	result := ""

	for _, v := range s {
		if IsFirst {
			IsFirst = false
			if v >= 'a' && v <= 'z' {
				result += string(v-32)
			} else {
				result += string(v)
			}
		} else {
			if v >= 'A' && v <= 'Z' {
				result += string(v+32)
			} else {
				result += string(v)
			}
		}
	}
	return result
}
