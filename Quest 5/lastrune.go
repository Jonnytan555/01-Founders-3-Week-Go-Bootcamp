package piscine

func LastRune(s string) rune {
	aStringChanged := []rune(s)

	l := len(aStringChanged)

	return aStringChanged[l-1]
}