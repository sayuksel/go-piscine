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
// <<<<<<< HEAD
			} else if v >= '0' && v <= '9'{
				result += string(v)
			} else {
				IsFirst = true
// =======
// >>>>>>> f5489433b08036113dd6f56e0ba170a8ba3bfd77
				result += string(v)
			}
		}
	}
	return result
}
