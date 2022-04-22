package piscine

func isPower(nb int) bool {
	if nb == 0 {
		return false
	}

	if nb == 1 {
		return true
	}

	if nb > 0 {
		result := 1

		for i := 1; result <= nb; i++ {
			result = result * 2
			if result == nb {
				return true
			}
		}
		return false
	} else {
		return false
	}
}
