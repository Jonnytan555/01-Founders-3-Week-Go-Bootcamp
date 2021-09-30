package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	num := 0

	for _, char := range a {
		if char > num {
			num = char
		}
	}
	return num
}
