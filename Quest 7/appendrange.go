package piscine

// Write a function that takes an int min and an int max as parameters

// That function returns a slice of int with all the values between min and max

func AppendRange(min, max int) []int {
	var answer []int

	for i := min - 1; i < max-1; i++ {
		answer = append(answer, i+1)
	}

	return answer
}
