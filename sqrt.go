package piscine

func Sqrt(nb int) int {
	prod := 0

	for i := 2; i < nb; i++ {
		prod = i * i
		if prod == nb {
			return i
		}
	}
	return 0
}
