package piscine

func ConcatParams(args []string) string {
	s := ""

	for i := 0; i < len(args); i++ {
		s += args[i]
		s += "\n"
	}
	return s
}
