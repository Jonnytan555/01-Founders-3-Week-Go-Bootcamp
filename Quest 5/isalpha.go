package piscine

func IsAlpha(s string) bool {
	for _, char := range s {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') && (char < 48 || char > 57) {
			return false
		}
	}
	return true
}
