package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// store arguments inside a variable, don't include file reference
	arguments := os.Args[1:]

	// loop through the arguments
	for i := range arguments {
		// loop through the characters in the arguments, printing out characters
		for _, val := range arguments[i] {
			z01.PrintRune(val)
		}
		z01.PrintRune('\n')
	}
}
