package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	numb := ""
	if n == 0 {
		z01.PrintRune('0')
	}
	neg := false

	if n < 0 {
		neg = true
	}

	if neg {
		n = n * -1
	}

	// toPrint := "-9223372036854775808"

	// for _, char := range toPrint {
	// 	z01.PrintRune(char)
	// }
	// 	if n == - {
	// 		// z01.PrintRune('-')
	// 		n = 922337203685477580
	// 	numb += "8"
	// }

	for i := 0; n > 0; i++ {
		numb = string(rune((n%10)+48)) + numb
		n = n / 10

	}

	if neg {
		z01.PrintRune('-')
	}

	for _, char := range numb {
		z01.PrintRune(char)
	}
}
