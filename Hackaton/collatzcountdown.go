package piscine

// Then each term is obtained from the previous term as follows:

// if the previous term is even, the next term is one half of the previous term.

// If the previous term is odd, the next term is 3 times the previous term plus 1.
// The conjecture is that no matter what value of n, the sequence will always reach 1.

func CollatzCountdown(start int) int {
	count := 0

	if start <= 0 {
		return -1
	}

	for numb := start; numb != 1; {
		if numb%2 == 0 {
			numb = numb / 2
		} else {
			numb = (3 * numb) + 1
		}
		count++
	}
	return count
}
