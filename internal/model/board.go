package model

import "strings"

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

func (b Board) InBounds(row, col int) bool {
	return row >= 0 && row < b.Size && col >= 0 && col < b.Size
}

func (b Board) Get(row, col int) rune {
	return b.Cells[row][col]
}

func (b *Board) Set(row, col int, value rune) {
	b.Cells[row][col] = value
}

func (b Board) Clone() Board {
	cloned := Board{
		Size:  b.Size,
		Cells: make([][]rune, b.Size),
	}

	for row := 0; row < b.Size; row++ {
		cloned.Cells[row] = make([]rune, b.Size)
		copy(cloned.Cells[row], b.Cells[row])
	}

	return cloned
}

func (b Board) String() string {
	rows := make([]string, b.Size)

	for row := 0; row < b.Size; row++ {
		rows[row] = string(b.Cells[row])
	}

	return strings.Join(rows, "\n")
}
