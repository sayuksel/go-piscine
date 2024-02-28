package piscine

func Capitalize(s string) string {

	var result string
	wordStart := true

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			if wordStart {
				if char >= 'a' && char <= 'z' {
					result += string(char - 32)
				} else {
					result += string(char)
				}
				wordStart = false
			} else {
				if char >= 'A' && char <= 'Z' {
					result += string(char + 32)
				} else {
					result += string(char)
				}
			}
		} else {
			if char >= 'A' && char <= 'Z' {
				result += string(char + 32)
			} else if char >= '0' && char <= '9' {
				result += string(char)
			} else {
				wordStart = true

			}
		}
	}
	return result

	// first := true
	// result := ""

	// for _, v := range s {
	// 	if first && (v >= 'a' && v <= 'z') {
	// 		first = false
	// 		result += string(v - 32)
	// 	} else if !first {
	// 		if v >= 'A' && v <= 'Z' {
	// 			result += string(v + 32)
	// 		} else if v >= 'a' && v <= 'z' {
	// 			result += string(v)
	// 		} else {
	// 			first = true
	// 			result += string(v)
	// 		}
	// 	}
	// }
	// return result

	// 	IsFirst := true
	// 	result := ""

	// 	for _, v := range s {
	// 		if IsFirst {
	// 			IsFirst = false
	// 			if v >= 'a' && v <= 'z' {
	// 				result += string(v-32)
	// 			} else {
	// 				result += string(v)
	// 			}
	// 		} else {
	// 			if v >= 'A' && v <= 'Z' {
	// 				result += string(v+32)
	// <<<<<<< HEAD
	// 			} else if v >= '0' && v <= '9'{
	// 				result += string(v)
	// 			} else {
	// 				IsFirst = true
	// =======
	// 			} else {
	// >>>>>>> f5489433b08036113dd6f56e0ba170a8ba3bfd77
	// 				result += string(v)
	// 			}
	// 		}
	// 	}
	// 	return result
}
