package piscine

func Compact(ptr *[]string) int {
	var a []string

	for _, char := range *ptr {
		if char == "" || char == "false" || char == "0" {
			continue
		} else {
			a = append(a, char)
		}
	}

	*ptr = a

	return len(a)
}
