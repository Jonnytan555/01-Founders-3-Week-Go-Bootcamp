package piscine

func Rot14(s string) string {
	aString := ""

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			if char > 'L' {
				aString = aString + string(char-12)
			} else {
				aString = aString + string(char+14)
			}
		} else if char >= 'a' && char <= 'z' {
			if char > 'l' {
				aString = aString + string(char-12)
			} else {
				aString = aString + string(char+14)
			}
		} else {
			aString = aString + string(char)
		}
	}
	return aString
}
