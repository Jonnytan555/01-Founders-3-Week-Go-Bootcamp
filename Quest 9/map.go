package piscine

func Map(f func(int) bool, a []int) []bool {
	var answer []bool

	for _, char := range a {
		answer = append(answer, f(char))
	}
	return answer
}
