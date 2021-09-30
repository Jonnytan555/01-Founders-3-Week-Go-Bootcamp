package main

import "github.com/01-edu/z01"

type point struct {
	x string
	y string
}

func setPoint(points *point) {
	points.x = "42"
	points.y = "21"
}

func main() {
	points := &point{}

	setPoint(points)

	aString := "x = " + points.x + ", " + "y = " + points.y

	for _, char := range aString {
		z01.PrintRune(char)
	}
}
