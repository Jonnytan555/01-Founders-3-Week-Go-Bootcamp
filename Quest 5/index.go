package piscine

func Index(s string, toFind string) int {
	for i := range s {
		if len(toFind) < len(s[i:]) {
			if string(s[i:i+len(toFind)]) == toFind {
				return i
			}
		}
	}
	return -1
}
