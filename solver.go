package sudoku

import (
	"fmt"
)

func Solve(board *[][]int) {

	for v := 0; v < 9; v++ {
		for h := 0; h < 9; h++ {
			if (*board)[v][h] == 0 {
				for n := 0; n <= 9; n++ {
					if IsPossible(v, h, n, board) {
						(*board)[v][h] = n
						// If n can be placed, recall the solve function
						Solve(board)
						// If after the recursive call is made and a number cannot be placed
						// Then a 0 is placed and the whole solving process will start again
						(*board)[v][h] = 0
					}
				}
				// Once all of the numbers are placed function will terminate
				return
			}
		}
	}

	for i, line := range *board {
		for j := range line {
			if j != len(line)-1 {
				fmt.Printf("%v ", (*board)[i][j])
			} else {
				fmt.Printf("%v", (*board)[i][j])
			}

		}
		fmt.Println()
	}
}
