package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var numb []rune

	if n == 0 {
		numb = append(numb, 48)
	}

	for i := 0; n > 0; i++ {
		numb = append(numb, rune(n%10)+48)
		n = n / 10
	}

	for i := range numb {
		if i == len(numb)-1 {
		} else {
			for j := 0; j < len(numb)-1; j++ {
				if numb[j] > numb[j+1] {
					numb[j], numb[j+1] = numb[j+1], numb[j]
				}
			}
		}
	}

	for _, char := range numb {
		z01.PrintRune(char)
	}
}
