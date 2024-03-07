package piscine

func CountIf(f func(string) bool, tab []string) int {
	count := 0

	for _, v := range tab {
		retval := f(v)

		if retval {
			count++
		}
	}
	return count
}
