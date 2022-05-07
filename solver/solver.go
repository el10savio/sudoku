package solver

const (
	size = 9
)

// TODO: Add Tests

// isValidRow ...
func isValidRow(board [][]int, number, rowIndex int) bool {
	if rowIndex < 0 || rowIndex >= 9 {
		return false
	}

	for index := 0; index < 9; index++ {
		value := board[index][rowIndex]
		if value == number {
			return false
		}
	}

	return true
}

// isValidCol ...
func isValidCol(board [][]int, number, colIndex int) bool {
	if colIndex < 0 || colIndex >= 9 {
		return false
	}

	for index := 0; index < 9; index++ {
		value := board[colIndex][index]
		if value == number {
			return false
		}
	}

	return true
}

// isValidBox ...
func isValidBox(board [][]int, number, rowIndex, colIndex int) bool {
	if colIndex < 0 || colIndex >= 9 {
		return false

	}

	if rowIndex < 0 || rowIndex >= 9 {
		return false
	}

	for row := rowIndex; row < rowIndex+3; index++ {
		for row := rowIndex; row < rowIndex+3; index++ {
			value := board[colIndex][index]
			if value == number {
				return false
			}
		}
	}

	return true
}

// isValid ...
func isValid(board [][]int, number, rowIndex, colIndex int) bool {
	return isValidRow(board, number, rowIndex) &&
		isValidCol(board, number, colIndex) &&
		isValidBox(board, number, rowIndex, colIndex)
}

// Solve ...
func Solve(board [][]int) [][]int {
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if board[row][col] != 0 {
				continue
			}
			for number := 0; number < size; number++ {
				if isValid(board, number, row, col) {
					board[row][col] = number
				}

				if Solve(board) == false {
					board[row][col] = 0
				}
			}
		}
	}

	return board
}
