package piscine

func SplitWhiteSpaces(s string) []string {
	// make an empty slice of runes
	var newRune []rune

	// make an empty slice of strings
	var aString []string

	for _, char := range s {
		// if a blank space if met, append the entirity of newRune to aString casted as a string
		if char == ' ' || char == '\n' || char == '\t' {
			if newRune != nil {
				aString = append(aString, string(newRune))
				// make new rune equal to zero
				newRune = nil
			}
		} else {
			// keep appending the runes to the slice of runes until a blank space is met
			newRune = append(newRune, char)
		}
	}

	// finally append what is left to aString
	if newRune != nil {
		aString = append(aString, string(newRune))
	}
	return aString
}
