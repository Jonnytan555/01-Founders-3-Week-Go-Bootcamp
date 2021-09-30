package piscine

// Write a function that returns True for a string slice

// if when the string slice is passed through the f function, at least one elemnt returns true

func Any(f func(string) bool, a []string) bool {
	var answer []bool

	for _, char := range a {
		answer = append(answer, f(char))
	}

	for _, char := range answer {
		if char {
			return true
		}
	}
	return false
}
