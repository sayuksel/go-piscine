package piscine

func TrimAtoi(s string) int {

	num := 0
	IsNegative := false
	IsNumber := false

	for _, char := range s {
		if char >= '0' && char <= '9' {
			num = num * 10 + int(char) - '0'
			IsNumber = true 
		} else if char == '-' && IsNumber == false{
			IsNegative = true
		} 
	}

	if IsNegative{ 
		num *= -1
	}

	return num
	// str_arr := []rune(s)
	// //int_arr := []int{}
	// var ret_val int
	// flag := false

	// for _, v := range str_arr {
	// 	if v >= '0' && v <= '9' {
	// 		ret_val = append(ret_val, int(v) + 48)
	// 		flag = true
	// 	} else if v == 45 && flag == false {
	// 		ret_val = append(ret_val, '-')
	// 	}
	// }
	// //ret_val = int_arr
	// return ret_val
}
