package piscine

func Map(f func(int) bool, a []int) []bool {
	arr := []bool{}

	for _, v := range a {
		retval := f(v)
		arr = append(arr, retval)
	}
	return arr
}
