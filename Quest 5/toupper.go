package piscine

func ToUpper(s string) string {
	aRune := []rune(s)

	for i, char := range aRune {
		if char >= 'a' && char <= 'z' {
			aRune[i] = aRune[i] - 32
		}
	}
	return string(aRune)
}
