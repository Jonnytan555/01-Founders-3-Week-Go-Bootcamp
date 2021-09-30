package piscine

// Write a function IsSorted that returns true, if the slice of int is sorted, and that returns false otherwise.

// The function passed in parameter returns a positive int if a (the first argument) is superior to b (the second argument), it returns 0 if they are equal and it returns a negative int otherwise.

// To do your testing you have to write your own f function.

func Sort(a int, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := range a {
		if i < len(a)-2 {
			if f(a[i], a[i+1]) < 0 {
				if f(a[i+1], a[i+2]) > 0 {
					return false
				}
			} else if f(a[i], a[i+1]) > 0 {
				if f(a[i+1], a[i+2]) < 0 {
					return false
				}
			}
		}
	}
	return true
}
