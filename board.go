package sudoku

func Board(board *[][]int, input *[]string) {
	// build an array of 9 arrays, each storing 9 empty elements
	*board = make([][]int, 9)

	for i := range *board {
		(*board)[i] = make([]int, 9)
	}

	// Fill up the board with the input arguments
	for i, line := range *input {
		for j, element := range line {
			// convert the string numbers to integers
			if element != '.' {
				atoi := int(element - 48)
				// convert ascii to integer for every j element of every line i
				(*board)[i][j] = atoi
			}
		}
	}
}
