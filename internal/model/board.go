package model

const EmptyCell = '.'

type Board struct {
	Size  int
	Cells [][]rune
}

func NewBoard(size int) Board {
	board := Board{
		Size:  size,
		Cells: make([][]rune, size),
	}

	for row := 0; row < size; row++ {
		board.Cells[row] = make([]rune, size)
		for col := 0; col < size; col++ {
			board.Cells[row][col] = EmptyCell
		}
	}

	return board
}
