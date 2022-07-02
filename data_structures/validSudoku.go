// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
//
// * Each row must contain the digits 1-9 without repetition.
// * Each column must contain the digits 1-9 without repetition.
// * Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
//
// Note:
// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.
//
// Examples: see validSudoku_test.go

package datastructures

import "math"

func isValidSudoku(board [][]byte) bool {
	n := len(board)
	m := int(math.Sqrt(float64(n))) // m x m is the size of subBoxes
	boxesSets := make([]uint16, n)
	columnsSets := make([]uint16, n)

	// row traversal
	for i := 0; i < n; i++ {
		var rowSet uint16
		for j := 0; j < n; j++ {
			digit := board[i][j]
			if digit == '.' {
				continue
			}

			boxI := i / m
			boxJ := j / m
			boxSeqNumber := m*boxI + boxJ
			if containsOrInsert(&rowSet, digit) ||
				containsOrInsert(&columnsSets[j], digit) ||
				containsOrInsert(&boxesSets[boxSeqNumber], digit) {
				return false
			}
		}
	}
	return true
}

func containsOrInsert(set *uint16, digit byte) bool {
	num := digit - '0'
	if *set&(1<<num) != 0 {
		return true
	}
	*set = *set | (1 << num)
	return false
}
