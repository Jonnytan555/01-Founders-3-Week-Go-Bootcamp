package piscine

func ToLower(s string) string {
	word := ""

	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			word = word + string(char+32)
		} else {
			word = word + string(char)
		}
	}
	return word
}
