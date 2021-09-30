package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// store arguments after filename in a variable
	arguments := os.Args[1:]

	// find the length of variable
	larg := len(arguments)

	// iterate from the final argument element
	for i := larg - 1; i >= 0; i-- {
		// print out the characters
		for _, val := range arguments[i] {
			z01.PrintRune(val)
		}
		z01.PrintRune('\n')
	}
}
