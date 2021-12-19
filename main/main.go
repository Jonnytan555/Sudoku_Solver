package main

import (
	"fmt"
	"os"
	"sudoku"
)

var board [][]int
var Input = os.Args[1:]

func main() {
	if sudoku.Check(Input) {
		sudoku.Board(&board, &Input)
		sudoku.IsPossible(2, 0, 5, &board)
		sudoku.Solve(&board)
	} else {
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("Error - Incorrect input format")
		fmt.Println("--------------------------------------------------------------")
		fmt.Println()
		fmt.Println("1. Must be 9 digits per line")
		fmt.Println("2. Only numbers and .'s allowed")
		fmt.Println("3. Must not be less than 17 digits revealed")

	}
}
