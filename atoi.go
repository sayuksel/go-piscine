package piscine

/*func Atoi(s string) int {
	result := 0
	ascii := 0
	for i := 0; i < len(s); i++ {
		ascii = int(s[i])
		if ascii == 43 || ascii == 45 {
			result = (result * 10) + int(s[i]-48)
			if ascii == 43{
				result[0] =
			}
		}
}
*/

func Atoi(s string) int {
	posflag := false
	negflag := false
	result := 0
	ascii := 0
	for i := 0; i < len(s); i++ {
		ascii = int(s[i])
		if ascii >= 48 && ascii <= 57 {
			result = (result * 10) + (ascii - 48)
		} else if (ascii == 45 || ascii == 43) && (negflag || posflag ){
			return 0
		} else if ascii == 45 {
			negflag = true
		}else if ascii == 43 {
			posflag = true
		} else {
			return 0
		}
	}

	if negflag {
		result = result * -1
	}
	return result
}
