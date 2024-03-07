package piscine

func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} 
	return -1
}


func IsSorted(f func(a, b int) int, a []int) bool {
	l := len(a)
	flag := true

	for i := 1; i < l; i++ {
		if f(a[i-1], a[i]) == 1 && flag == true {
			flag = true
		} else {
			flag = false
		}
	}
	return flag
}
