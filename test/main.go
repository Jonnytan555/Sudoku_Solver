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
		fmt.Print("Error")
		
		/*
		fmt.Println("--------------------------------------------------------------")
		fmt.Println("Error - Incorrect input format")
		fmt.Println("--------------------------------------------------------------")
		fmt.Println()
		fmt.Println("1. Must be 9 digits per line")
		fmt.Println("2. Only numbers and .'s allowed")
		fmt.Println("3. Must not be less than 17 digits revealed")
		fmt.Println()
		fmt.Println("Example: go run . .96.4...1 1...6...4 5.481.39. ..795..43 .3..8.... 4.5.23.18 .1.63..59 .59.7.83. ..359...7 ")
		*/
	}
		
}