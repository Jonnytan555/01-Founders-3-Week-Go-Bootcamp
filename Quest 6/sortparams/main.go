package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguements := os.Args[1:]

	for i := 0; i < len(arguements); i++ {
		for j := range arguements {
			if j+1 < len(arguements) {
				if arguements[j] > arguements[j+1] {
					arguements[j], arguements[j+1] = arguements[j+1], arguements[j]
				}
			}
		}
	}
	for _, v := range arguements {
		for _, arg := range v {
			z01.PrintRune(arg)
		}
		z01.PrintRune('\n')
	}
}
