package piscine

func Capitalize(s string) string {
	IsFirst := true
	result := ""

	for _, v := range s {
		if IsFirst {
			if v >= 'a' && v <= 'z' {
				result += string(v-32)
			} else {
				result += string(v)
			}
			IsFirst = false
		} else {
			result += string(v)
		}	
	}
	return result
}
