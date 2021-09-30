package piscine

func NRune(s string, n int) rune {
	sChanged := []rune(s)

	if n <= 0 {
		return 0
	}

	l := len(sChanged)

	if n <= l {
		return sChanged[n-1]
	} else {
		return 0
	}
}
