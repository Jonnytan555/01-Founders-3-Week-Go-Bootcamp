package main

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := 1; j <= 8; j++ {
			for k := 2; k <= 9; k++ {
				if i == 7 && j == 8 && k == 9 {
					z01.PrintRune(rune(i) + 48)
					z01.PrintRune(rune(j) + 48)
					z01.PrintRune(rune(k) + 48)
					z01.PrintRune('\n')
				} else if i < j && j < k {
					z01.PrintRune(rune(i) + 48)
					z01.PrintRune(rune(j) + 48)
					z01.PrintRune(rune(k) + 48)
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
}
