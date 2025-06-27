package piscine

import (
	"github.com/01-edu/z01"
)

const boardSize = 8

// SolveEightQueens recursively solves the eight queens puzzle.
func SolveEightQueens(board [boardSize]int, row int) {
	if row == boardSize {
		for _, chislo := range board {
			z01.PrintRune(rune(chislo + 49))
		}
		z01.PrintRune(10)
		// fmt.Println(board)
		return
	}

	for col := 0; col < boardSize; col++ {
		if isSafe(board, row, col) {
			board[row] = col
			SolveEightQueens(board, row+1)
			board[row] = 0
		}
	}
}

// isSafe checks if a queen can be placed in a given row and column without
// threatening any other queens.
func isSafe(board [boardSize]int, row int, col int) bool {
	for r := 0; r < row; r++ {
		c := board[r]
		if c == col || c-row == col-r || c+row == col+r {
			return false
		}
	}
	return true
}

func EightQueens() {
	var board [boardSize]int
	SolveEightQueens(board, 0)
}
