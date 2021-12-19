package sudoku

func Check(arguments []string) bool {
	// args length not equal to 9

	if len(arguments) != 9 {
		return false
	}

	// testing for only numbers or .'s

	for i := range arguments {
		for _, char := range arguments[i] {
			if char >= '1' && char <= '9' || char == '.' {
			} else {
				return false
			}
		}
	}

	// number of revealed pieces must be 17 or more
	count := 0

	for i := range arguments {
		for _, char := range arguments[i] {
			if char >= '1' && char <= '9' {
				count++
			}
		}
	}

	if count >= 17 {
		return true
	} else {
		return false
	}
}