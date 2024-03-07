package piscine

func f(a, b int) int {
	if a < b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	retval := f(a[0], a[1])

	if retval > 0 {
		return true
	} else {
		return false
	}
}
