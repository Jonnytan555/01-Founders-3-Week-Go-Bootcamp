package piscine

func CountIf(f func(string) bool, tab []string) int {
	// var answer []bool

	count := 0

	for i := range tab {
		if f(tab[i]) {
			// answer = append(answer, f(tab[i]))
			count++
		}
	}
	return count
}
