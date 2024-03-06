package piscine

func Any(f func(string) bool, a []string) bool {
	for _, v := range a {
		retval := f(v)

		if retval {
			return true
		}
	}
	return false
}
