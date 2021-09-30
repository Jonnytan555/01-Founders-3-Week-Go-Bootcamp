package piscine

func BasicJoin(elems []string) string {
	aString := ""

	for _, char := range elems {
		aString = aString + char
	}
	return aString
}
