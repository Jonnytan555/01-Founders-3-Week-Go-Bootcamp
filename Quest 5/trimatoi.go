package piscine

func TrimAtoi(s string) int {
	st := []rune(s)

	numbs := ""
	neg := false

	for _, char := range s {
		if char >= 48 && char <= 57 {
			numbs = numbs + string(char)
		}
	}

	var b int = 0

	for _, char := range numbs {
		c := int(char) - 48
		b = (b * 10) + c
	}

	for i, char := range s {
		if char == 45 {
			for j := 0; j < i; j++ {
				if st[j] >= 48 && st[j] <= 57 {
					neg = false
				} else {
					neg = true
				}
			}
		}
	}
	if neg {
		b = b * (-1)
	}
	return b
}
