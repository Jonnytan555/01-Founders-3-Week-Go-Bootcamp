package piscine

func ConcatParams(args []string) string {
	var s string

	if len(args) > 0 {
		for i, arg := range args {
			if i < len(args)-1 {
				s = s + string(arg)
				s = s + ("\n")
			} else {
				s = s + string(arg)
			}
		}
	}

	return s
}
