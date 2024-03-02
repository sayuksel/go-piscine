package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	arr := make([]int, max)

	for i := min; i < max; i++ {
		arr[i] = i
	}
	return arr[min:]
}
