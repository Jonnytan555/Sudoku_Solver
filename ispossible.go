package sudoku

func IsPossible(horozontal int, vertical int, input int, board *[][]int) bool {

	// Check Horozontally, vertically and in side a 3 by 3 box, if the input value fits
	for v := 0; v < 9; v++ {
		if (*board)[horozontal][v] == input {
			return false
		}

		for h := 0; h < 9; h++ {
			if (*board)[h][vertical] == input {
				return false
			}
		}

		// Go has built in floor divison, can now iterate between a moving window
		h0 := (horozontal / 3) * 3
		v0 := (vertical / 3) * 3

		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if (*board)[h0+i][v0+j] == input {
					return false
				}
			}
		}

	}
	return true
}
