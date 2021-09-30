package main

import "github.com/01-edu/z01"

func main() {
	PrintComb2()
}

func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for j := 1; j <= 99; j++ {
			if j > i {
				z01.PrintRune(rune(i/10 + 48))
				z01.PrintRune(rune(i%10 + 48))
				z01.PrintRune(32)
				z01.PrintRune(rune(j/10 + 48))
				z01.PrintRune(rune(j%10 + 48))
				if i != 98 || j != 99 {
					z01.PrintRune(44)
					z01.PrintRune(32)
				}
			}
		}
	}
}
