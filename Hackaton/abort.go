package piscine

func Abort(a, b, c, d, e int) int {
	s := []int{a, b, c, d, e}

	for i := 0; i < len(s)-1; i++ {
		for i := range s {
			if i < len(s)-1 {
				if s[i] < s[i+1] {
					s[i], s[i+1] = s[i+1], s[i]
				}
			}
		}
	}

	result := s[2]

	return result
}
