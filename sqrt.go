package piscine

func Sqrt(nb int) int {
	prod := 0

	if nb == 1 {
		return 1
	}

	if nb == 0 {
		return 0
	}

	for i := 0; i < nb; i++ {
		prod = i * i
		if prod == nb {
			return i
		}
	}
	return 0
}
