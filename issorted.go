package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	n := len(a)
	if n == 1 { // if didn't work swap the flags from false to true
		return true
	}
	asce := false
	desc := false
	for i := 1; i < n; i++ {
		if f(a[i-1], a[i]) <= 0 {
			asce = true
		} else {
			desc = true
		}
	}
	if asce == desc {
		return false
	} else {
		return true
	}
}
