package piscine

func IsNumeric(s string) bool {
	for _, char := range s {
		if char < 48 || char > 57 {
			return false
		}
	}
	return true
}
