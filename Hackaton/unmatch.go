package piscine

func Unmatch(a []int) int {
	count := 0
	var num int

	for i := 0; i < len(a); i++ {
		for _, char := range a {
			if i == char {
				count++
			}
		}
		if count%2 != 0 {
			num = i
			return num
		}
		count = 0
	}
	return -1
}
