package piscine

func Swap(a *int, b *int) {
	a_temp := *a
	b_temp := *b

	*a = b_temp
	*b = a_temp
}
