package piscine

func ConcatParams(args []string) string {
	s := ""
	l := len(args)

	for i := 0; i < l; i++ {
		s += args[i]
		if i < l-1 {
			s += "\n"
		}
	}
	return s
}
