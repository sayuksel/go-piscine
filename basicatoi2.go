package piscine

func BasicAtoi2(s string) int {
	result := 0
	ascii := 0
	for i := 0; i < len(s); i++ {
		ascii = int(s[i])
		if ascii >= 48 && ascii <= 57 {
			result = (result * 10) + (ascii - 48)
		} else {
			return 0
		}

	}
	return result
}
