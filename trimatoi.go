package piscine

func TrimAtoi(s string) int {
	num := 0
	IsNegative := false
	IsNumber := false

	for _, v := range s {
		if v >= '0' && v <= '9' {
			num = num*10 + int(v) - '0'
			IsNumber = true
		} else if v == '-' && !IsNumber {
			IsNegative = true
		}
	}
	if IsNegative {
		num *= -1
	}
	return num
}
