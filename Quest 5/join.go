package piscine

func Join(strs []string, sep string) string {
	aString := ""

	for i, char := range strs {
		if i != len(strs)-1 {
			aString = aString + char + sep
		} else {
			aString = aString + char
		}
	}
	return aString
}
