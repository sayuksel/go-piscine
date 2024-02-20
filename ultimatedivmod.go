package piscine

func UltimateDivMod(a *int, b *int) {
	a_temp := *a
	*a = *a / *b
	*b = a_temp % *b
}
