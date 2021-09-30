package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// create a new variable, select the first element, slice the ./
	argument := os.Args[0][2:]

	// print out the individual characters
	for _, char := range argument {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}
